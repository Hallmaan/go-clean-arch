package transport_http_transaction

import (
	transaction_controller "clean_arch_ws/internal/controller/transaction"
	transaction_domain "clean_arch_ws/internal/entities/transaction"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type OutputResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type TransactionTransporter struct {
	TransactionCreateController transaction_controller.CreateTransactionControllerPorts
}

func NewTransactionTransporter(trxController transaction_controller.CreateTransactionControllerPorts) *TransactionTransporter {
	return &TransactionTransporter{
		TransactionCreateController: trxController,
	}
}

func (t *TransactionTransporter) CreateTransaction(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	domain := transaction_domain.TransactionDomain{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&domain)

	if err != nil {
		fmt.Println("error nya ini ", err)
	}

	handler, err := t.TransactionCreateController.Create(&domain)

	if err != nil {
		fmt.Println("error transporter", err)
	}

	resp := OutputResponse{
		Message: "Berhasil",
		Data:    handler,
	}

	w.Header().Add("Content-type", "application/json")

	encoder := json.NewEncoder(w)
	err = encoder.Encode(resp)
}
