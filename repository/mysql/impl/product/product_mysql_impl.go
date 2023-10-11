package product_mysql_impl

import (
	"clean_arch_ws/internal/entities"
	"clean_arch_ws/repository/initializer"
	mysql_ports "clean_arch_ws/repository/mysql/ports"
	"context"
	"fmt"
)

type ProductMysqlRepositoryImpl struct {
	conn *initializer.Replication
}

func NewProductMysqlRepositoryImpl(conn *initializer.Replication) mysql_ports.ProductRepository {
	return &ProductMysqlRepositoryImpl{
		conn: conn,
	}
}

func (pd ProductMysqlRepositoryImpl) GetProduct(ctx context.Context, id int) (*entities.ProductDomain, error) {
	sql := fmt.Sprintf(`select * from products where id = %d`, id)

	pdDomain := &entities.ProductDomain{}

	err := pd.conn.Standby.GetContext(ctx, pdDomain, sql)

	if err != nil {
		return nil, err
	}

	return pdDomain, nil
}
