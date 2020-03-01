// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package main

import (
	"os"
	"os/signal"
	"strings"

	"github.com/mattermost/mattermost-bot-sample-golang/config"

	"github.com/mattermost/mattermost-server/v5/model"
)

var client *model.Client4
var webSocketClient *model.WebSocketClient

var botUser *model.User
var botTeam *model.Team
var debuggingChannel *model.Channel

// Documentation for the Go driver can be found
// at https://godoc.org/github.com/mattermost/platform/model#Client
func main() {
	println(config.SampleName)

	SetupGracefulShutdown()

	client = model.NewAPIv4Client("http://localhost:8065")

	// Lets test to see if the mattermost server is up and running
	MakeSureServerIsRunning()

	// lets attempt to login to the Mattermost server as the bot user
	// This will set the token required for all future calls
	// You can get this token with client.AuthToken
	LoginAsTheBotUser()

	// If the bot user doesn't have the correct information lets update his profile
	UpdateTheBotUserIfNeeded()

	// Lets find our bot team
	FindBotTeam()

	// This is an important step.  Lets make sure we use the botTeam
	// for all future web service requests that require a team.
	//client.SetTeamId(botTeam.Id)

	// Lets create a bot channel for logging debug messages into
	CreateBotDebuggingChannelIfNeeded()
	SendMsgToDebuggingChannel("_"+config.SampleName+" has **started** running_", "")

	// Lets start listening to some channels via the websocket!
	webSocketClient, err := model.NewWebSocketClient4("ws://localhost:8065", client.AuthToken)
	if err != nil {
		println("We failed to connect to the web socket")
		PrintError(err)
	}

	webSocketClient.Listen()

	go func() {
		for {
			select {
			case resp := <-webSocketClient.EventChannel:
				HandleWebSocketResponse(resp)
			}
		}
	}()

	// You can block forever with
	select {}
}

// MakeSureServerIsRunning ...
func MakeSureServerIsRunning() {
	if props, resp := client.GetOldClientConfig(""); resp.Error != nil {
		println("There was a problem pinging the Mattermost server.  Are you sure it's running?")
		PrintError(resp.Error)
		os.Exit(1)
	} else {
		println("Server detected and is running version " + props["Version"])
	}
}

// LoginAsTheBotUser ...
func LoginAsTheBotUser() {
	if user, resp := client.Login(config.UserEmail, config.UserPassword); resp.Error != nil {
		println("There was a problem logging into the Mattermost server.  Are you sure ran the setup steps from the README.md?")
		PrintError(resp.Error)
		os.Exit(1)
	} else {
		botUser = user
	}
}

//UpdateTheBotUserIfNeeded ..
func UpdateTheBotUserIfNeeded() {
	if botUser.FirstName != config.UserFirstName || botUser.LastName != config.UserLastName || botUser.Username != config.Username {
		botUser.FirstName = config.UserFirstName
		botUser.LastName = config.UserLastName
		botUser.Username = config.Username

		if user, resp := client.UpdateUser(botUser); resp.Error != nil {
			println("We failed to update the Sample Bot user")
			PrintError(resp.Error)
			os.Exit(1)
		} else {
			botUser = user
			println("Looks like this might be the first run so we've updated the bots account settings")
		}
	}
}

//FindBotTeam ..
func FindBotTeam() {
	if team, resp := client.GetTeamByName(config.TeamName, ""); resp.Error != nil {
		println("We failed to get the initial load")
		println("or we do not appear to be a member of the team '" + config.TeamName + "'")
		PrintError(resp.Error)
		os.Exit(1)
	} else {
		botTeam = team
	}
}

//CreateBotDebuggingChannelIfNeeded ..
func CreateBotDebuggingChannelIfNeeded() {
	if rchannel, resp := client.GetChannelByName(config.ChannelLogName, botTeam.Id, ""); resp.Error != nil {
		println("We failed to get the channels")
		PrintError(resp.Error)
	} else {
		debuggingChannel = rchannel
		return
	}

	// Looks like we need to create the logging channel
	channel := &model.Channel{}
	channel.Name = config.ChannelLogName
	channel.DisplayName = "Debugging For Sample Bot"
	channel.Purpose = "This is used as a test channel for logging bot debug messages"
	channel.Type = model.CHANNEL_OPEN
	channel.TeamId = botTeam.Id
	if rchannel, resp := client.CreateChannel(channel); resp.Error != nil {
		println("We failed to create the channel " + config.ChannelLogName)
		PrintError(resp.Error)
	} else {
		debuggingChannel = rchannel
		println("Looks like this might be the first run so we've created the channel " + config.ChannelLogName)
	}
}

//SendMsgToDebuggingChannel ..
func SendMsgToDebuggingChannel(msg string, replyToID string) {
	post := &model.Post{}
	post.ChannelId = debuggingChannel.Id
	post.Message = msg

	post.RootId = replyToID

	if _, resp := client.CreatePost(post); resp.Error != nil {
		println("We failed to send a message to the logging channel")
		PrintError(resp.Error)
	}
}

// SendMsgToAnyChannel ..
func SendMsgToAnyChannel(msg string, replyToID string, channelID string) {
	post := &model.Post{}
	post.ChannelId = channelID
	post.Message = msg

	post.RootId = replyToID

	if _, resp := client.CreatePost(post); resp.Error != nil {
		println("We failed to send a message to the specified channel")
		PrintError(resp.Error)
	}
}

//HandleWebSocketResponse ..
func HandleWebSocketResponse(event *model.WebSocketEvent) {
	HandleMsgFromDebuggingChannel(event)
}

//HandleMsgFromDebuggingChannel ..
func HandleMsgFromDebuggingChannel(event *model.WebSocketEvent) {

	// Lets only reponded to messaged posted events
	if event.Event != model.WEBSOCKET_EVENT_POSTED {
		return
	}

	println("responding to debugging channel msg")

	post := model.PostFromJson(strings.NewReader(event.Data["post"].(string)))
	if post != nil {
		// ignore my events
		if post.UserId == botUser.Id {
			return
		}

		if msg, match := publishItEverywhere(post); match {
			SendMsgToAnyChannel(msg, post.Id, post.ChannelId)
			return
		}
	}

}

// PrintError .
func PrintError(err *model.AppError) {
	println("\tError Details:")
	println("\t\t" + err.Message)
	println("\t\t" + err.Id)
	println("\t\t" + err.DetailedError)
}

// SetupGracefulShutdown ..
func SetupGracefulShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			if webSocketClient != nil {
				webSocketClient.Close()
			}

			SendMsgToDebuggingChannel("_"+config.SampleName+" has **stopped** running_", "")
			os.Exit(0)
		}
	}()
}
