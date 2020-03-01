package linkedin

// GetLinkedinAuthLink gets the linkedin authorization link for a given channel
func GetLinkedinAuthLink(channelID string) string {
	return linkedinConfig.AuthCodeURL(channelID)
}
