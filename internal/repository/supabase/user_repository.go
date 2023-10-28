package supabase

import (
	"fmt"
	"github.com/nedpals/supabase-go"
	"github.com/timopattikawa/jubelio-chatapp/domain"
	"github.com/timopattikawa/jubelio-chatapp/dto"
	"log"
)

type UserRepository struct {
	Supabase *supabase.Client
}

func (u UserRepository) GetUserByUsername(username string) (domain.User, error) {
	var results []map[string]interface{}
	err := u.Supabase.DB.From("users").Select("*").
		Eq("username", username).
		Execute(&results)

	if err != nil {
		log.Fatal("Fail in get by username Supabase", err)
		return domain.User{}, err
	}

	if len(results) == 0 {
		return domain.User{}, fmt.Errorf("this is an %s error", "user not found")
	}

	theID := uint(results[0]["id"].(float64))
	if theID == 0 {
		log.Println("Some thing wrong when get id")
		return domain.User{}, fmt.Errorf("this is an %s error", "convert id")
	}
	unPackage := domain.User{
		Id:       theID,
		Username: fmt.Sprintf("%v", results[0]["username"]),
		Password: fmt.Sprintf("%v", results[0]["password"]),
	}

	return unPackage, nil
}

func (u UserRepository) SaveUser(user dto.AuthReq) error {
	var results []domain.User
	err := u.Supabase.DB.From("users").Insert(user).Execute(&results)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func NewUserRepository(conn *supabase.Client) domain.UserRepository {
	return &UserRepository{
		Supabase: conn,
	}
}
