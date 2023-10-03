package transaction_ucase_ports

import (
	transaction_domain "clean_arch_ws/pkg/domain/transaction"
	"context"
)

type CreateTransactionUseCasePorts interface {
	Create(ctx context.Context, trx transaction_domain.TransactionDomain) (*transaction_domain.TransactionDomain, error)
	// create with product dll
}
