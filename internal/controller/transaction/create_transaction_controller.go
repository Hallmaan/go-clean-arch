package transaction_controller

import (
	"clean_arch_ws/internal/entities"
	ucase_transaction "clean_arch_ws/internal/usecase/transaction"
	mysql_ports "clean_arch_ws/repository/mysql/ports"
	"context"
	"fmt"
)

type CreateTransactionControllerPorts interface {
	Create(*entities.TransactionDomain) (*entities.TransactionDomain, error)
}

type CreateTransactionController struct {
	CreateTransactionUseCase ucase_transaction.CreateTransactionUseCasePorts
	TransactionRepository    mysql_ports.TransactionRepository
}

func NewCreateTransactionController(ucase ucase_transaction.CreateTransactionUseCasePorts, repo mysql_ports.TransactionRepository) CreateTransactionControllerPorts {
	return &CreateTransactionController{
		CreateTransactionUseCase: ucase,
		TransactionRepository:    repo,
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
