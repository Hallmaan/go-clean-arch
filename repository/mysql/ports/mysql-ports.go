package mysql_ports

import (
	"clean_arch_ws/internal/entities"
	"context"
)

type TransactionRepository interface {
	Get(ctx context.Context, transaction_id int64) (*entities.TransactionDomain, error)
	Create(ctx context.Context, transaction *entities.TransactionDomain) (int64, error)
}

type ProductRepository interface {
	Get(ctx context.Context, id int) (*entities.ProductDomain, error)
}
