package ucase_product

import (
	"clean_arch_ws/internal/entities"
	mysql_ports "clean_arch_ws/repository/mysql/ports"
	"context"
	"errors"
)

type GetProductUcasePorts interface {
	Get(ctx context.Context, id int) (*entities.ProductDomain, error)
	// get by id, name dll
}

type GetProductUcase struct {
	ProductRepository mysql_ports.ProductRepository
}

func NewGetProductByIdUCase(pdRepo mysql_ports.ProductRepository) GetProductUcasePorts {
	return &GetProductUcase{
		ProductRepository: pdRepo,
	}
}

func (pd GetProductUcase) Get(ctx context.Context, id int) (*entities.ProductDomain, error) {
	res, err := pd.ProductRepository.Get(ctx, id)

	if res == nil || err != nil {
		return nil, errors.New("product not found")
	}

	return res, nil
}
