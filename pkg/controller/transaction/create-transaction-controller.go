package transaction_controller

import (
	transaction_controller_ports "clean_arch_ws/pkg/controller/transaction/ports"
	transaction_domain "clean_arch_ws/pkg/domain/transaction"
	transaction_ucase_ports "clean_arch_ws/pkg/usecase/transaction/ports"
	mysql_ports "clean_arch_ws/repository/mysql/ports"
	"context"
	"fmt"
)

type CreateTransactionController struct {
	CreateTransactionUseCase transaction_ucase_ports.CreateTransactionUseCasePorts
	TransactionRepository    mysql_ports.TransactionRepository
}

func NewCreateTransactionController(ucase transaction_ucase_ports.CreateTransactionUseCasePorts, repo mysql_ports.TransactionRepository) transaction_controller_ports.CreateTransactionControllerPorts {
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
