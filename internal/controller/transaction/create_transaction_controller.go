package transaction_controller

import (
	"clean_arch_ws/internal/entities"
	ucase_transaction "clean_arch_ws/internal/usecase/transaction"
	"context"
	"fmt"
)

type CreateTransactionControllerPorts interface {
	Create(*entities.TransactionDomain) (*entities.TransactionDomain, error)
}

type CreateTransactionController struct {
	CreateTransactionUseCase ucase_transaction.CreateTransactionUseCasePorts
}

func NewCreateTransactionController(ucase ucase_transaction.CreateTransactionUseCasePorts) CreateTransactionControllerPorts {
	return &CreateTransactionController{
		CreateTransactionUseCase: ucase,
	}
}

func (c CreateTransactionController) Create(trx *entities.TransactionDomain) (*entities.TransactionDomain, error) {
	ctx := context.Background()
	x, err := c.CreateTransactionUseCase.Create(ctx, trx)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// LISTEN TO NATS

	return x, nil
}
