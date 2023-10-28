package domain

import (
	"github.com/timopattikawa/jubelio-chatapp/dto"
)

type User struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRepository interface {
	GetUserByUsername(username string) (User, error)
	SaveUser(user dto.AuthReq) error
}

type UserService interface {
	RegistrationUser(req dto.AuthReq) (string, error)
	AuthUser(req dto.AuthReq) (dto.AuthRes, error)
}
