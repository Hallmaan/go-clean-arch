package websocket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/nats-io/nats.go"
)

type WebSocketTransports struct {
	NatsJs nats.JetStreamContext
}

func NewWSTransports(nats nats.JetStreamContext) *WebSocketTransports {
	return &WebSocketTransports{
		NatsJs: nats,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Memeriksa origin permintaan (untuk keamanan, biasanya Anda harus mengatur kebijakan yang tepat)
		return true
	},
}

func (ws WebSocketTransports) HandleWebSocket(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client terhubung ke WebSocket!")

	for {
		// Membaca pesan dari klien
		// messageType, p, err := conn.ReadMessage()
		// if err != nil {
		// 	fmt.Println("Error:", err)
		// 	return
		// }

		_, err = ws.NatsJs.Subscribe("TransactionUpdates", func(msg *nats.Msg) {
			// Kirim data yang diterima ke klien WebSocket
			data := msg.Data
			// Kirim data ke klien WebSocket yang terhubung
			// Menampilkan pesan yang diterima di server
			fmt.Printf("Menerima pesan: %s\n", data)

			// Mengirim pesan balasan kembali ke klien
			err = conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		})
	}
}
