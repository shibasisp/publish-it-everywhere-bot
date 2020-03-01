package main

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-bot-sample-golang/publisher"
	"github.com/mattermost/mattermost-bot-sample-golang/publisher/linkedin"
	"github.com/mattermost/mattermost-bot-sample-golang/publisher/twitter"
	"github.com/mattermost/mattermost-bot-sample-golang/types"
	"github.com/mattermost/mattermost-server/v5/model"
)

func publishItEverywhere(post *model.Post) (string, bool) {

	if post.Message == ":authenticate_linkedin" {
		linkedinAuthLink := linkedin.GetLinkedinAuthLink(post.ChannelId)
		msg := fmt.Sprintf("Go to the given URL to authenticate with linkedin: %s", linkedinAuthLink)
		return msg, true
	}
	if post.Message == ":authenticate_twitter" {
		resp, err := twitter.Login(post)
		if err != nil {
			fmt.Println(err)
			return "Couldn't fetch the Auth URL", false
		}
		msg := fmt.Sprintf("Go to the given URL to authenticate with twitter: %s", resp)
		return msg, true
	}

	if strings.HasPrefix(post.Message, ":publish_twitter") {
		message := trimFirstWord(post.Message)
		return publishToTwitter(message, post)
	}
	if strings.HasPrefix(post.Message, ":publish_linkedin") {
		message := trimFirstWord(post.Message)
		return publishToLinkedIn(message, post)
	}

	if strings.HasPrefix(post.Message, ":publish_it_everywhere") {
		message := trimFirstWord(post.Message)
		_, postedLinkedIn := publishToLinkedIn(message, post)
		_, postedTwitter := publishToTwitter(message, post)

		if postedLinkedIn && postedTwitter {
			return "Successfully posted on Twitter and LinkedIn", true
		}
		if postedLinkedIn && !postedTwitter {
			return "Successfully posted on LinkedIn but couldn't post on Twitter, Please login again on twitter", true
		}
		if !postedLinkedIn && postedTwitter {
			return "Successfully posted on Twitter but couldn't post on LinkedIn, Please login again on LinkedIn", true
		}
		if !postedLinkedIn && postedTwitter {
			return "Couldn't post on Twitter and LinkedIn", true
		}
	}

	return "", false
}

func trimFirstWord(inp string) string {
	return strings.Join(strings.Split(inp, " ")[1:], " ")
}
func publishToTwitter(message string, post *model.Post) (string, bool) {
	msg, err := publisher.Do(types.TwitterPlatform, post.ChannelId, message)
	if err != nil {
		return "", false
	}
	return msg, true
}
func publishToLinkedIn(message string, post *model.Post) (string, bool) {
	msg, err := publisher.Do(types.LinkedinPlatform, post.ChannelId, message)
	if err != nil {
		return "", false
	}
	return msg, true
}
