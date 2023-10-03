package product_controller_ports

import (
	product_domain "clean_arch_ws/pkg/domain/product"
)

type GetProductControllerPorts interface {
	Get(id int) (*product_domain.ProductDomain, error)
	// get by id, name dll
}
