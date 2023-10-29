package psql

import (
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/timopattikawa/jubelio-chatapp/domain"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

func TestSaveMessagePsql_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Fail Create sql Mock", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	chat := domain.Chat{
		ID:       1,
		Sender:   1,
		Receiver: 2,
		Message:  "Hi Bro!",
		CreateAt: fmt.Sprintf("%s", time.Now()),
	}

	mock.ExpectExec("INSERT INTO chat").WithArgs(
		chat.ID,
		chat.Sender,
		chat.Receiver,
		chat.Message,
		chat.CreateAt,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	psql := NewChatRepositoryPsql(db)
	_, err = psql.SaveMessagePsql(chat)
	if err != nil {
		t.Errorf("error '%s' was not expected, while inserting a row", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindAllChatById_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "sender", "receiver", "message", "created_at"}).
		AddRow(1, 1, 3, "Hi Bro", time.Now()).
		AddRow(2, 1, 3, "Bro", time.Now())

	query := "SELECT id, sender, receiver, message, created_at FROM chat WHERE receiver = \\$1 AND sender = \\$2"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := NewChatRepositoryPsql(db)

	anArticle, err := a.FindAllChatById(1, 3)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}
