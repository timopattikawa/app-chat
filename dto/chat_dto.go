package dto

type ChatDto struct {
	Sender   uint   `json:"sender"`
	Receiver uint   `json:"receiver"`
	Message  string `json:"message"`
}
