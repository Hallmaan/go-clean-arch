package transaction_controller_ports

import (
	transaction_domain "clean_arch_ws/pkg/domain/transaction"
)

type CreateTransactionControllerPorts interface {
	Create(transaction_domain.TransactionDomain) (*transaction_domain.TransactionDomain, error)
	// get by id, name dll
}
