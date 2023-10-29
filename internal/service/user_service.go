package service

import (
	"errors"
	"github.com/timopattikawa/jubelio-chatapp/domain"
	"github.com/timopattikawa/jubelio-chatapp/dto"
	"github.com/timopattikawa/jubelio-chatapp/internal/security"
	"log"
)

type UserService struct {
	repository domain.UserRepository
}

func (u UserService) RegistrationUser(req dto.AuthReq) (string, error) {
	userCheck, _ := u.repository.GetUserByUsername(req.Username)
	if userCheck.Username != "" {
		log.Println("user has been register")
		return "", errors.New("fail to register user because user has been register")
	}
	err := u.repository.SaveUser(req)
	if err != nil {
		log.Println("Fail To save user internal server error", err)
		return "", err
	}
	return "register success", nil
}

func (u UserService) AuthUser(req dto.AuthReq) (dto.AuthRes, error) {
	userData, err := u.repository.GetUserByUsername(req.Username)
	if err != nil {
		log.Println("Nof Found User")
		return dto.AuthRes{}, err
	}
	if userData.Password != req.Password {
		return dto.AuthRes{}, errors.New("wrong password")
	}

	jwt, err := security.GenerateJwt(userData.Id, userData.Username)
	if err != nil {
		log.Println("Fail to generate jwt")
		return dto.AuthRes{}, err
	}

	return dto.AuthRes{
		Token: string(jwt),
	}, nil
}

func NewUserService(userRepository domain.UserRepository) domain.UserService {
	return &UserService{
		userRepository,
	}
}
