package router

import (
	transporter "clean_arch_ws/transports/http"
	"clean_arch_ws/transports/websocket"

	"github.com/julienschmidt/httprouter"
	"github.com/nats-io/nats.go"
)

func NewRouter(
	transactionTransporter *transporter.TransactionTransporter,
	nats *nats.Conn,
) *httprouter.Router {
	router := httprouter.New()

	ws := websocket.NewWSTransports(nats)
	router.GET("/ws", ws.HandleWebSocket)
	router.POST("/transaction", transactionTransporter.CreateTransaction)
	return router
}
