package domain

import (
	"github.com/timopattikawa/jubelio-chatapp/dto"
)

type Chat struct {
	ID       uint   `json:"id"`
	Sender   uint   `json:"sender"`
	Receiver uint   `json:"receiver"`
	Message  string `json:"message"`
	CreateAt string `json:"create_at"`
}

type ChatService interface {
	SendMessage(chatDto dto.ChatDto) (dto.ChatDto, error)
	FetchLastMessage(receiver uint) (dto.ChatDto, error)
	SearchHistoryByReceiver(senderId uint, receiver uint) ([]string, error)
}

type ChatRepository interface {
	SaveMessage(model dto.ChatDto) (Chat, error)
	SaveMessagePsql(model Chat) (Chat, error)
	GetChatById(chatId uint) (Chat, error)
	GetChatByIdReceiver(receiver uint) (Chat, error)
	FindAllChatById(senderId uint, receiver uint) ([]Chat, error)
}
