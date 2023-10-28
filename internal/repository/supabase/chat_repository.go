package supabase

import (
	"fmt"
	"github.com/nedpals/supabase-go"
	"github.com/timopattikawa/jubelio-chatapp/domain"
	"github.com/timopattikawa/jubelio-chatapp/dto"
	"log"
	"time"
)

type ChatRepository struct {
	sClient *supabase.Client
}

func (c ChatRepository) SaveMessagePsql(model domain.Chat) (domain.Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c ChatRepository) GetChatByIdReceiver(receiver uint) (domain.Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c ChatRepository) SaveMessage(model dto.ChatDto) (domain.Chat, error) {
	var results []domain.Chat
	err := c.sClient.DB.From("chat").Insert(model).Execute(&results)
	log.Println(results)
	if err != nil {
		log.Println(err)
		return domain.Chat{}, err
	}
	return results[0], err
}

func (c ChatRepository) GetChatById(id uint) (domain.Chat, error) {
	var results []map[string]interface{}
	err := c.sClient.DB.From("chat").Select("*").
		Eq("id", fmt.Sprintf("%v", id)).
		Execute(&results)

	if err != nil {
		log.Fatal("Fail in get by username Supabase", err)
		return domain.Chat{}, err
	}

	if len(results) == 0 {
		return domain.Chat{}, fmt.Errorf("this is an %s error", "user not found")
	}

	log.Printf("Get data chat by id : %v", results)
	date := fmt.Sprintf("%v", results[0]["create_at"])
	log.Println(time.Parse(time.DateTime, date))

	theID := uint(results[0]["id"].(float64))
	theSender := uint(results[0]["sender"].(float64))
	theReceiver := uint(results[0]["receiver"].(float64))
	if theID == 0 || theSender == 0 || theReceiver == 0 {
		log.Println("Some thing wrong when get id")
		return domain.Chat{}, fmt.Errorf("this is an %s error", "convert id")
	}

	unPackage := domain.Chat{
		ID:       theID,
		Sender:   theSender,
		Receiver: theReceiver,
		Message:  fmt.Sprintf("%v", results[0]["message"]),
	}

	return unPackage, nil
}

func (c ChatRepository) FindAllChatById(senderId uint, receiver uint) ([]domain.Chat, error) {
	//TODO implement me
	panic("implement me")
}

func NewChatRepositorySup(client *supabase.Client) domain.ChatRepository {
	return &ChatRepository{
		sClient: client,
	}
}
