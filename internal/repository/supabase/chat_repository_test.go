package supabase

import (
	supa "github.com/nedpals/supabase-go"
	"github.com/stretchr/testify/assert"
	"github.com/timopattikawa/jubelio-chatapp/domain"
	"github.com/timopattikawa/jubelio-chatapp/dto"
	"testing"
)

func TestChatRepository_GetChatById(t *testing.T) {
	sClient := supa.CreateClient(url, key)
	expectedData := domain.Chat{
		ID:       1,
		Sender:   1,
		Receiver: 3,
		Message:  "Hi bro",
	}
	sup := NewChatRepositorySup(sClient)
	data, err := sup.GetChatById(1)
	if err != nil {
		t.Fatal("Fail to get chat by id", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedData.ID, data.ID)
	assert.Equal(t, expectedData.Sender, data.Sender)
	assert.Equal(t, expectedData.Receiver, data.Receiver)
}

func TestChatRepository_SaveMessage(t *testing.T) {
	sClient := supa.CreateClient(url, key)

	data := dto.ChatDto{
		Sender:   1,
		Receiver: 3,
		Message:  "Ping",
	}
	sup := NewChatRepositorySup(sClient)
	message, err := sup.SaveMessage(data)
	if err != nil {
		t.Fatal("Fail to save chat", err)
	}

	dataChat, err := sup.GetChatById(message.ID)
	if err != nil {
		t.Fatal("Fail to get chat by id", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, dataChat.ID, message.ID)
	assert.Equal(t, dataChat.Message, message.Message)
	assert.Equal(t, dataChat.Sender, message.Sender)
	assert.Equal(t, dataChat.Receiver, message.Receiver)
}
