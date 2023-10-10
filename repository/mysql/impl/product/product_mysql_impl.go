package product_mysql_impl

import (
	"clean_arch_ws/internal/entities"
	mysql_ports "clean_arch_ws/repository/mysql/ports"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProductMysqlRepositoryImpl struct {
	conn *sqlx.DB
}

func NewProductMysqlRepositoryImpl(conn *sqlx.DB) mysql_ports.ProductRepository {
	return &ProductMysqlRepositoryImpl{
		conn: conn,
	}
}

func (pd ProductMysqlRepositoryImpl) Get(ctx context.Context, id int) (*entities.ProductDomain, error) {
	sql := fmt.Sprintf(`select * from products where id = %d`, id)

	pdDomain := &entities.ProductDomain{}

	err := pd.conn.GetContext(ctx, pdDomain, sql)

	if err != nil {
		return nil, err
	}

	return pdDomain, nil
}
