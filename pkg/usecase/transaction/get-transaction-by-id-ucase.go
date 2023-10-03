package ucase_transaction

import (
	transaction_domain "clean_arch_ws/pkg/domain/transaction"
	product_ucase_ports "clean_arch_ws/pkg/usecase/product/ports"
	mysql_ports "clean_arch_ws/repository/mysql/ports"
	"context"

	"github.com/jmoiron/sqlx"
)

type GetTransactionByIdUCase struct {
	TransactionRepo mysql_ports.TransactionRepository
	GetProductUcase product_ucase_ports.GetProductUcasePorts
	connectionDb    *sqlx.DB
}

func NewGetTransactionByIdUCase(trxRepo mysql_ports.TransactionRepository, conn *sqlx.DB, pdUcase product_ucase_ports.GetProductUcasePorts) *GetTransactionByIdUCase {
	return &GetTransactionByIdUCase{
		TransactionRepo: trxRepo,
		GetProductUcase: pdUcase,
		connectionDb:    conn,
	}
}

func (trx AddNewTransactionUCase) GetTransactionById(ctx context.Context, id int64) (*transaction_domain.TransactionDomain, error) {
	trxCreate, err := trx.TransactionRepo.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	pd, err := trx.GetProductUcase.Get(ctx, trxCreate.Product.ID)

	if err != nil {
		return nil, err
	}

	trxDomain, err := transaction_domain.NewTransaction(trxCreate.GetId(), trxCreate.GetName(), pd)

	return trxDomain, nil
}
