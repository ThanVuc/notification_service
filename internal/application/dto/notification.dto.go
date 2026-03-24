package dto

type WorkGenerationNotificationMessage struct {
	ID              *string  `json:"id,omitempty" bson:"_id,omitempty"`
	Title           string   `json:"title" bson:"title"`
	Message         string   `json:"message" bson:"message"`
	Link            *string  `json:"link,omitempty" bson:"link,omitempty"`
	SenderID        string   `json:"sender_id" bson:"sender_id"`
	ReceiverIDs     []string `json:"receiver_ids" bson:"receiver_ids"`
	ImageURL        *string  `json:"img_url,omitempty" bson:"img_url,omitempty"`
	CorrelationID   string   `json:"correlation_id" bson:"correlation_id"`
	CorrelationType int      `json:"correlation_type" bson:"correlation_type"`
}

type TeamNotificationMessage struct {
	EventType   string                          `json:"event_type" bson:"event_type"`
	SenderID    string                          `json:"sender_id" bson:"sender_id"`
	ReceiverIDs []string                        `json:"receiver_ids" bson:"receiver_ids"`
	Payload     TeamNotificationMessagePayload  `json:"payload" bson:"payload"`
	Metadata    TeamNotificationMessageMetadata `json:"metadata" bson:"metadata"`
}

type TeamNotificationMessagePayload struct {
	Title           string  `json:"title" bson:"title"`
	Message         string  `json:"message" bson:"message"`
	Link            *string `json:"link,omitempty" bson:"link,omitempty"`
	ImageURL        *string `json:"img_url,omitempty" bson:"img_url,omitempty"`
	CorrelationID   string  `json:"correlation_id" bson:"correlation_id"`
	CorrelationType int     `json:"correlation_type" bson:"correlation_type"`
}

type TeamNotificationMessageMetadata struct {
	IsSentMail           bool     `json:"is_sent_mail" bson:"is_sent_mail"`
	NonExistentReceivers []string `json:"non_existent_receivers,omitempty" bson:"non_existent_receivers,omitempty"`
}
