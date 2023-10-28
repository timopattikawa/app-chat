package psql

//
//import (
//	"database/sql"
//	"github.com/timopattikawa/jubelio-chatapp/domain"
//	"log"
//)
//
//type UserRepository struct {
//	db *sql.DB
//}
//
//func (u UserRepository) GetUserByUsername(username string) (domain.User, error) {
//	var data = domain.User{}
//	err := u.db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).
//		Scan(&data.Id, &data.Username, &data.Password)
//	if err != nil {
//		return domain.User{}, err
//	}
//	return data, err
//}
//
//func (u UserRepository) SaveUser(user domain.User) error {
//	query := "INSERT INTO users (user, password) VALUES (? , ?)"
//	_, err := u.db.Exec(query, user.Username, user.Password)
//	if err != nil {
//		log.Fatal("Fail to save user", err)
//		return err
//	}
//	return err
//}
//
//func NewUserRepoPsql(con *sql.DB) domain.UserRepository {
//	return &UserRepository{
//		db: con,
//	}
//}
