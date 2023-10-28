package domain

import (
	"github.com/gorilla/websocket"
	"log"
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

func GetDataFromSupabaseRealtime() *websocket.Conn {
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
		log.Println(err)
		return nil
	}

	//done := make(chan struct{})

	return conn
}
