package supabase

import (
	supa "github.com/nedpals/supabase-go"
	"github.com/stretchr/testify/assert"
	"github.com/timopattikawa/jubelio-chatapp/dto"
	"log"
	"testing"
)

func TestUserRepository_GetUserByUsername(t *testing.T) {
	supabaseClient := supa.CreateClient(url, key)

	userRepository := NewUserRepository(supabaseClient)
	user, err := userRepository.GetUserByUsername("timopattikawa")
	if err != nil {
		log.Fatal("Testing fail get username")
		return
	}
	assert.Equal(t, user.Username, "timopattikawa")
	assert.Equal(t, user.Id, uint(1))
	assert.Equal(t, user.Password, "asdfasdf")
}

func TestUserRepository_GetUserByUsername_ButNull(t *testing.T) {
	supabaseClient := supa.CreateClient(url, key)

	userRepository := NewUserRepository(supabaseClient)
	_, err := userRepository.GetUserByUsername("asdfasdf")
	assert.Error(t, err)
}

func TestUserRepository_SaveUser(t *testing.T) {
	supabaseClient := supa.CreateClient(url, key)

	userRepository := NewUserRepository(supabaseClient)
	data := dto.AuthReq{
		Username: "usercoba",
		Password: "jagobanget",
	}
	err := userRepository.SaveUser(data)
	if err != nil {
		log.Fatal("Testing fail get username")
		return
	}

	user, err := userRepository.GetUserByUsername(data.Username)
	if err != nil {
		log.Fatal("Testing fail get username")
		return
	}
	assert.Equal(t, user.Username, data.Username)
}
