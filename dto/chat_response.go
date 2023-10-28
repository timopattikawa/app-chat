package dto

type ChatRes struct {
	Sender   int      `json:"Sender"`
	Receiver int      `json:"receiver"`
	Messages []string `json:"messages"`
}
