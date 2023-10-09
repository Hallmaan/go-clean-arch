package ucase_transaction

import (
	transaction_domain "clean_arch_ws/internal/entities/transaction"
	ucase_product "clean_arch_ws/internal/usecase/product"
	mysql_ports "clean_arch_ws/repository/mysql/ports"
	"context"

	"github.com/jmoiron/sqlx"
)

type GetTransactionByIdUCase struct {
	TransactionRepo mysql_ports.TransactionRepository
	GetProductUcase ucase_product.GetProductUcasePorts
	connectionDb    *sqlx.DB
}

func NewGetTransactionByIdUCase(trxRepo mysql_ports.TransactionRepository, conn *sqlx.DB, pdUcase ucase_product.GetProductUcasePorts) *GetTransactionByIdUCase {
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
