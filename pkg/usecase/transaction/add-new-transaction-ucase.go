package ucase_transaction

import (
	transaction_domain "clean_arch_ws/pkg/domain/transaction"
	product_ports "clean_arch_ws/pkg/usecase/product/ports"
	transaction_ucase_ports "clean_arch_ws/pkg/usecase/transaction/ports"
	mysql_ports "clean_arch_ws/repository/mysql/ports"
	"context"
	"errors"
)

type AddNewTransactionUCase struct {
	TransactionRepo mysql_ports.TransactionRepository
	GetProductUcase product_ports.GetProductUcasePorts
}

func NewAddTrxUsecase(trxRepo mysql_ports.TransactionRepository, pdUcase product_ports.GetProductUcasePorts) transaction_ucase_ports.CreateTransactionUseCasePorts {
	return &AddNewTransactionUCase{
		TransactionRepo: trxRepo,
		GetProductUcase: pdUcase,
	}
}

func (trx AddNewTransactionUCase) Create(ctx context.Context, p transaction_domain.TransactionDomain) (*transaction_domain.TransactionDomain, error) {
	res, err := trx.GetProductUcase.Get(ctx, p.Product.ID)

	if res == nil || err != nil {
		return nil, errors.New("product not found")
	}

	trxCreateId, err := trx.TransactionRepo.Create(ctx, &p)

	if err != nil {
		return nil, err
	}

	getTrx, err := trx.TransactionRepo.Get(ctx, trxCreateId)

	if err != nil {
		return nil, err
	}

	return getTrx, nil
}
