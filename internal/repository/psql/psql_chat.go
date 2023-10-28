package psql

import (
	"database/sql"
	"fmt"
	"github.com/timopattikawa/jubelio-chatapp/domain"
	"github.com/timopattikawa/jubelio-chatapp/dto"
	"log"
)

type ChatRepository struct {
	conn *sql.DB
}

func (c ChatRepository) SaveMessagePsql(model domain.Chat) (domain.Chat, error) {
	query := "INSERT INTO chat (id, sender, receiver, message, created_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := c.conn.Exec(query, model.ID, model.Sender, model.Receiver, model.Message, model.CreateAt)
	if err != nil {
		log.Fatal("Fail to save user", err)
		return domain.Chat{}, err
	}
	return domain.Chat{}, err
}

func (c ChatRepository) SaveMessage(model dto.ChatDto) (domain.Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c ChatRepository) GetChatById(chatId uint) (domain.Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c ChatRepository) GetChatByIdReceiver(receiver uint) (domain.Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c ChatRepository) FindAllChatById(senderId uint, receiver uint) ([]domain.Chat, error) {
	var data []domain.Chat
	rows, err := c.conn.Query(
		"SELECT id, sender, receiver, message, created_at FROM chat WHERE receiver = $1 AND sender = $2",
		receiver, senderId)
	if err != nil {
		log.Println("error select ", err)
		return []domain.Chat{}, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	for rows.Next() {
		var each = domain.Chat{}
		var err = rows.Scan(&each.ID, &each.Sender, &each.Receiver, &each.Message, &each.CreateAt)

		if err != nil {
			fmt.Println(err.Error())
			return []domain.Chat{}, err
		}

		data = append(data, each)
	}

	return data, err
}

func NewChatRepositoryPsql(conn *sql.DB) domain.ChatRepository {
	return &ChatRepository{
		conn: conn,
	}
}
