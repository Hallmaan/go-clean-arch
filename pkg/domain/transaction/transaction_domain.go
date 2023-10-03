package transaction_domain

import (
	pd "clean_arch_ws/pkg/domain/product"
	"errors"
)

type TransactionDomain struct {
	ID      int64
	TrxName string `db:"transaction_name"`
	Product *pd.ProductDomain
}

func NewTransaction(id int64, name string, product *pd.ProductDomain) (trx *TransactionDomain, err error) {
	if id == 0 {
		return nil, errors.New("error id tidak boleh kosong")
	}

	if name == "agung" {
		return nil, errors.New("nama tidak boleh agung")
	}

	return &TransactionDomain{
		ID:      id,
		TrxName: name,
		Product: product,
	}, nil
}

func (trx TransactionDomain) GetName() string {
	return trx.TrxName
}

func (trx TransactionDomain) GetId() int64 {
	return trx.ID
}

func (trx TransactionDomain) GetProduct() pd.ProductDomain {
	return *trx.Product
}
