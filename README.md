## Noted Pembuatan Aplikasi:
1. List Penggunaan REST:
    - Registrasi dan authentikasi User
    - Pengiriman chat dengan REST
    - search history chat by id
2. Database menggunakan supabase untuk chat dan data user
3. Postgres Local untuk menyimpan history chat

## Keamanan:
- menggunakan JWT
- durasi exp 24 jam

### END POINT:
- `/registration` => Body json{username: string, password: string}
- `/auth` => Body json{username: string, password: string}
- `/message/send` => Body json {Sender : uint, Receiver: uint, Message:  string}
- `/message/history/?receiver={id}&sender={id]`

```UNTUK WEBSOCKET DIGUNAKAN KETIKA ADA SEND MESSAGE PADA SERVICE DIMANA LANGSUNG DIMASUKAN KEITKA ADA INSERT CHAT BARU```

## Database
```
TABLE CHAT (
   ID INT,
   SENDER INT,
   RECEIVER INT,
   MESSAGE STRING,
   CREATE_AT TIMESTAMP
)
```

Libraries Used:
- go get github.com/gofiber/contrib/websocket
- go get github.com/overseedio/realtime-go
- go get github.com/stretchr/testify/assert
- go get github.com/golang-jwt/jwt
- go get "github.com/gorilla/websocket"
- go get github.com/lib/pq
- go get github.com/gofiber/jwt/v2
- go get gopkg.in/DATA-DOG/go-sqlmock.v1

