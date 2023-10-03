package product_ucase_ports

import (
	product_domain "clean_arch_ws/pkg/domain/product"
	"context"
)

type GetProductUcasePorts interface {
	Get(ctx context.Context, id int) (*product_domain.ProductDomain, error)
	// get by id, name dll
}
