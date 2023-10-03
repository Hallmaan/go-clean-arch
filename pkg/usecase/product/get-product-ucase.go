package ucase_product

import (
	product_domain "clean_arch_ws/pkg/domain/product"
	mysql_ports "clean_arch_ws/repository/mysql/ports"
	"context"
	"errors"
)

type GetProductUcase struct {
	ProductRepository mysql_ports.ProductRepository
}

func NewGetProductByIdUCase(pdRepo mysql_ports.ProductRepository) *GetProductUcase {
	return &GetProductUcase{
		ProductRepository: pdRepo,
	}
}

func (pd GetProductUcase) Get(ctx context.Context, id int) (*product_domain.ProductDomain, error) {
	res, err := pd.ProductRepository.Get(ctx, id)

	if res == nil || err != nil {
		return nil, errors.New("product not found")
	}

	return res, nil
}
