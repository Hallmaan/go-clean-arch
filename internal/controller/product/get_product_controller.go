package product_controller

import (
	product_domain "clean_arch_ws/internal/entities/product"
	ucase_product "clean_arch_ws/internal/usecase/product"
	mysql_ports "clean_arch_ws/repository/mysql/ports"
	"context"
	"fmt"
)

type GetProductControllerPorts interface {
	Get(id int) (*product_domain.ProductDomain, error)
}

type GetProductController struct {
	GetProductUcase ucase_product.GetProductUcasePorts
	ProductRepo     mysql_ports.ProductRepository
}

func NewGetProductController(ucase ucase_product.GetProductUcasePorts, repo mysql_ports.ProductRepository) GetProductControllerPorts {
	return &GetProductController{
		GetProductUcase: ucase,
		ProductRepo:     repo,
	}
}

func (c GetProductController) Get(id int) (*product_domain.ProductDomain, error) {
	ctx := context.Background()
	x, err := c.GetProductUcase.Get(ctx, id)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// LISTEN TO NATS

	return x, nil
}
