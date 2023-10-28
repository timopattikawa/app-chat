package service

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/timopattikawa/jubelio-chatapp/domain"
	"github.com/timopattikawa/jubelio-chatapp/dto"
	"log"
	"time"
)

const wsURL = "wss://mkwvhlizfipwufxmoayf.supabase.co/realtime/v1/websocket?vsn=1.0.0&apikey=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Im1rd3ZobGl6Zmlwd3VmeG1vYXlmIiwicm9sZSI6ImFub24iLCJpYXQiOjE2OTgyOTMxNDMsImV4cCI6MjAxMzg2OTE0M30.xR-RqnDkXrQg6O8sSaehHs5DRitatX4R1B5j5TJDvkk"

type PostgresChanges struct {
	Event  string `json:"event"`
	Schema string `json:"schema"`
	Table  string `json:"table"`
}

type Config struct {
	PostgresChanges []PostgresChanges `json:"postgres_changes"`
}

type Payload struct {
	Config Config `json:"config"`
}

type Message struct {
	Topic   string      `json:"topic"`
	Event   string      `json:"event"`
	Payload interface{} `json:"payload"`
	Ref     string      `json:"ref"`
}

type ChatService struct {
	supRepository  domain.ChatRepository
	psqlRepository domain.ChatRepository
}

func (c ChatService) SendMessage(chatDto dto.ChatDto) (dto.ChatDto, error) {
	log.Printf("CONNECTING TO %s", wsURL)
	conn, h, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("RESPONSE CONNECT WS %s", h)
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("ADA ERR")
		}
	}(conn)

	message := Message{
		Topic: "realtime:public:chat",
		Event: "phx_join",
		Payload: Payload{
			Config: Config{
				PostgresChanges: []PostgresChanges{
					{Event: "*", Schema: "public", Table: "chat"},
				},
			},
		},
		Ref: "",
	}

	err = conn.WriteJSON(message)
	if err != nil {
		log.Println(err.Error())
		return dto.ChatDto{}, err
	}

	if err != nil {
		log.Println(err.Error())
		return dto.ChatDto{}, err
	}

	ctx, cfunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cfunc()
	log.Println("masuk routine")
	//defer close(done)
	saved := false
	var result = dto.ChatDto{}
	for {
		select {
		case <-ctx.Done():
			break
		default:
			log.Println("masuk loop")
			var message map[string]map[string]map[string]map[string]interface{}
			if err := conn.ReadJSON(&message); err != nil {
				log.Println("message read error:", err)
			}
			log.Println("recv: ", message)
			record := message["payload"]["data"]["record"]
			if record != nil {
				id := uint(record["id"].(float64))
				receiver := uint(record["receiver"].(float64))
				sender := uint(record["sender"].(float64))
				message := fmt.Sprintf("%v", record["message"])
				createdAt := fmt.Sprintf("%v", record["created_at"])
				chat := domain.Chat{
					ID:       id,
					Sender:   receiver,
					Receiver: sender,
					Message:  message,
					CreateAt: createdAt,
				}
				_, err = c.psqlRepository.SaveMessagePsql(chat)
				if err != nil {
					return dto.ChatDto{}, err
				}
				break
			}

			if !saved {
				resultDataSave, err := c.supRepository.SaveMessage(chatDto)
				if err != nil {
					log.Println(err.Error())
					return dto.ChatDto{}, err
				}
				result = dto.ChatDto{
					Sender:   resultDataSave.Sender,
					Receiver: resultDataSave.Receiver,
					Message:  resultDataSave.Message,
				}
				saved = true
			}
		}
	}

	return result, nil
}

func (c ChatService) FetchLastMessage(receiver uint) (dto.ChatDto, error) {
	//TODO implement me
	panic("implement me")
}

func (c ChatService) SearchHistoryByReceiver(senderId uint, receiver uint) ([]string, error) {
	listChat, err := c.psqlRepository.FindAllChatById(senderId, receiver)
	if err != nil {
		return nil, err
	}
	log.Println(listChat)
	var tmpResult []string

	for _, chat := range listChat {
		tmpResult = append(tmpResult, chat.Message)
	}

	return tmpResult, nil
}

func (c ChatService) SearchHistoryByDate(date time.Time) ([]dto.ChatDto, error) {
	//TODO implement me
	panic("implement me")
}

func NewChatService(sup domain.ChatRepository,
	psql domain.ChatRepository) domain.ChatService {
	return &ChatService{
		supRepository:  sup,
		psqlRepository: psql,
	}
}
