package linkedin

import (
	"net/url"

	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/linkedin"
)

var linkedinConfig *oauth2.Config

func init() {
	u, _ := url.Parse(config.TestServerURL)
	u.Path = "/api/linkedin/authenticate"

	linkedinConfig = &oauth2.Config{
		ClientID:    config.LinkedinClientID,
		RedirectURL: u.String(),
		Scopes:      config.LinkedinScopes,
		Endpoint:    linkedin.Endpoint,
	}
}
