package types

// PublishRequest ...
type PublishRequest struct {
	ChannelID string `json:"channel_id" bson:"channel_id"`
	Message   string `json:"message" bson:"message"`
}
