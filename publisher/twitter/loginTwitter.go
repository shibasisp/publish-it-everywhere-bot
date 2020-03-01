package twitter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"github.com/mattermost/mattermost-bot-sample-golang/types"
	"github.com/mattermost/mattermost-server/v5/model"
)

//Login ...
func Login(post *model.Post) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/twitter/authurl?channel_id=", config.ServerURL) + post.ChannelId)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var response types.Response

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}
	fmt.Println(response)
	return response.Data, nil
}
