package main

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-bot-sample-golang/publisher"
	"github.com/mattermost/mattermost-bot-sample-golang/publisher/linkedin"
	"github.com/mattermost/mattermost-bot-sample-golang/types"
	"github.com/mattermost/mattermost-server/v5/model"
)

func publishItEverywhere(post *model.Post) (string, bool) {

	if post.Message == ":authenticate_linkedin" {
		linkedinAuthLink := linkedin.GetLinkedinAuthLink(post.ChannelId)
		msg := fmt.Sprintf("Go to the given URL to authenticate with linkedin: %s", linkedinAuthLink)
		return msg, true
	}

	if strings.HasPrefix(post.Message, ":publish_it_everywhere ") {
		message := strings.TrimPrefix(post.Message, ":publish_it_everywhere ")

		linkedinResponse := publisher.Do(types.LinkedinPlatform, post.ChannelId, message)
		// TODO: replace with twitter API call
		twitterResponse := "twitter response"

		templateString := "Linkedin Response:\n%s\n\nTwitter Response:\n%s"
		return fmt.Sprintf(templateString, linkedinResponse, twitterResponse), true
	}

	return "", false
}
