package product_controller

import (
	product_controller_ports "clean_arch_ws/pkg/controller/product/ports"
	product_domain "clean_arch_ws/pkg/domain/product"
	product_ucase_ports "clean_arch_ws/pkg/usecase/product/ports"
	mysql_ports "clean_arch_ws/repository/mysql/ports"
	"context"
	"fmt"
)

type GetProductController struct {
	GetProductUcase product_ucase_ports.GetProductUcasePorts
	ProductRepo     mysql_ports.ProductRepository
}

func NewGetProductController(ucase product_ucase_ports.GetProductUcasePorts, repo mysql_ports.ProductRepository) product_controller_ports.GetProductControllerPorts {
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
