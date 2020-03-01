package publisher

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"github.com/mattermost/mattermost-bot-sample-golang/types"
	"github.com/mattermost/mattermost-bot-sample-golang/utils"
	"gopkg.in/yaml.v2"
)

// Do ...
func Do(platform types.Platform, channelID string, message string) string {
	var resp interface{}
	if err := utils.PostJSON(http.DefaultClient, fmt.Sprintf("%s/api/%s/publish", config.ServerURL, platform.String()), types.JSON{
		"channel_id": channelID,
		"message":    message,
	}, &resp); err != nil {
		return fmt.Sprintf("Error: %s\n", err)
	}

	bs, err := yaml.Marshal(resp)
	if err != nil {
		return fmt.Sprintf("Error: %s\n", err)
	}
	return string(bs)
}
