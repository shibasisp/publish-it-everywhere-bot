package publisher

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"github.com/mattermost/mattermost-bot-sample-golang/types"
	"github.com/mattermost/mattermost-bot-sample-golang/utils"
)

// Do ...
func Do(platform types.Platform, channelID string, message string) (string, error) {
	var resp interface{}
	if err := utils.PostJSON(http.DefaultClient, fmt.Sprintf("%s/api/%s/publish", config.ServerURL, platform.String()), types.PublishRequest{
		ChannelID: channelID,
		Message:   message,
	}, &resp); err != nil {
		return "", err
	}

	return fmt.Sprintf("Successfully posted to %s", platform), nil
}
