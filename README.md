```
Noted Penting Pak: 
Pada bagian /message/send enpoint service dapat melakukan save 
pada supabase lalu websocket dapat mengambil data dan menyimpan di
local DB. Tetapi request masih ada bug yaitu tidak mau berhenti atau
memberikan response balikan pada saat request 

Terima Kasih
```

Noted Pembuatan Aplikasi:
1. List Penggunaan REST:
    - Registrasi dan authentikasi User
    - Pengiriman chat dengan REST
    - search history chat by id
2. Database menggunakan supabase untuk chat dan data user
3. Postgres Local untuk menyimpan history chat

Keamanan:
- menggunakan JWT
- durasi exp 24 jam

END POINT:
- `/registration` => Body json{username: string, password: string}
- `/auth` => Body json{username: string, password: string}
- `/message/send` => Body json {Sender : uint, Receiver: uint, Message:  string}
- `/message/history/?receiver={id}&sender={id]`

```UNTUK WEBSOCKET DIGUNAKAN KETIKA ADA INSERT PADA SERVIEC DIMANA LANGSUNG DIMASUKAN KEITKA ADA INSERT CHAT BARU```

Libraries Used:
- go get github.com/gofiber/contrib/websocket
- go get github.com/overseedio/realtime-go
- go get github.com/stretchr/testify/assert
- go get github.com/golang-jwt/jwt 

