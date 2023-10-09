package transaction_controller

import (
	transaction_domain "clean_arch_ws/internal/entities/transaction"
	ucase_transaction "clean_arch_ws/internal/usecase/transaction"
	mysql_ports "clean_arch_ws/repository/mysql/ports"
	"context"
	"fmt"
)

type CreateTransactionControllerPorts interface {
	Create(*transaction_domain.TransactionDomain) (*transaction_domain.TransactionDomain, error)
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

func (c CreateTransactionController) Create(trx *transaction_domain.TransactionDomain) (*transaction_domain.TransactionDomain, error) {
	ctx := context.Background()
	x, err := c.CreateTransactionUseCase.Create(ctx, trx)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// LISTEN TO NATS

	return x, nil
}
