package structure

type Message struct {
	MessageId string `json:"message_id"`
	From      *User  `json:"from"`
	Chat      *Chat  `json:"chat"`
}
