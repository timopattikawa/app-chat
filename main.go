package main

import (
	"github.com/gofiber/fiber/v2"
	supa "github.com/nedpals/supabase-go"
	"github.com/timopattikawa/jubelio-chatapp/internal/config"
	"github.com/timopattikawa/jubelio-chatapp/internal/dbconfig"
	"github.com/timopattikawa/jubelio-chatapp/internal/handler"
	"github.com/timopattikawa/jubelio-chatapp/internal/repository/psql"
	"github.com/timopattikawa/jubelio-chatapp/internal/repository/supabase"
	"github.com/timopattikawa/jubelio-chatapp/internal/service"
	"log"
)

func main() {
	cfg := config.Get()
	supabaseClient := supa.CreateClient(cfg.SupabaseCon.Url, cfg.SupabaseCon.Key)
	chatRepositoryPSQL := dbconfig.PSQLGetConnection(cfg)

	userRepository := supabase.NewUserRepository(supabaseClient)
	userService := service.NewUserService(userRepository)

	chatRepository := supabase.NewChatRepositorySup(supabaseClient)
	repositoryPsql := psql.NewChatRepositoryPsql(chatRepositoryPSQL)
	chatService := service.NewChatService(chatRepository, repositoryPsql)

	app := fiber.New()
	handler.NewUserHandler(app, userService)
	handler.NewChatHandler(app, chatService)
	err := app.Listen(":9000")
	if err != nil {
		log.Fatal(err)
	}
}
