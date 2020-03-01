package twitter

import (
	"log"

	"github.com/mattermost/mattermost-bot-sample-golang/publisher"
	"github.com/mattermost/mattermost-bot-sample-golang/types"
	"github.com/mattermost/mattermost-server/v5/model"
)

//Publish ..
func Publish(post *model.Post) (bool, error) {

	resp, err := publisher.Do(types.TwitterPlatform, post.ChannelId, post.Message)
	if err != nil {
		return false, err
	}
	log.Println("Twitter publish response:", resp)

	return true, nil
}
