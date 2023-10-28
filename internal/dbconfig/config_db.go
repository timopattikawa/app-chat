package dbconfig

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nedpals/supabase-go"
	"github.com/spf13/viper"
	"github.com/timopattikawa/jubelio-chatapp/internal/config"
	"log"
)

func GetSupabaseConnection(config *config.Config) *supabase.Client {
	url := viper.GetString("URL_SUPABASE")
	key := viper.GetString("API_KEY_SUPABASE")
	supabase := supabase.CreateClient(url, key)

	return supabase
}

func PSQLGetConnection(cfg *config.Config) *sql.DB {
	dns := fmt.Sprintf(
		"host=%s "+
			"port=%s "+
			"user=%s "+
			"password=%s "+
			"dbname=%s "+
			"sslmode=disable ",
		cfg.DatabaseCon.Host,
		cfg.DatabaseCon.Port,
		"postgres",
		"postgres",
		cfg.DatabaseCon.Name)
	con, err := sql.Open("postgres", dns)

	if err != nil {
		log.Fatalf("Fail to open database %s", err.Error())
	}

	err = con.Ping()
	if err != nil {
		panic(err.Error())
	}

	return con
}
