package mysql_ports

import (
	product_domain "clean_arch_ws/pkg/domain/product"
	transaction_domain "clean_arch_ws/pkg/domain/transaction"
	"context"
)

type TransactionRepository interface {
	Get(ctx context.Context, transaction_id int64) (*transaction_domain.TransactionDomain, error)
	Create(ctx context.Context, transaction *transaction_domain.TransactionDomain) (int64, error)
}

type ProductRepository interface {
	Get(ctx context.Context, id int) (*product_domain.ProductDomain, error)
}
