package mysql

import (
	"clean_arch_ws/repository/initializer"
	"clean_arch_ws/repository/mysql/executor"
	product_mysql_impl "clean_arch_ws/repository/mysql/impl/product"
	transaction_mysql_impl "clean_arch_ws/repository/mysql/impl/transaction"
	mysql_ports "clean_arch_ws/repository/mysql/ports"
	"context"
)

type SqlInterface interface {
	mysql_ports.TransactionRepository
	mysql_ports.ProductRepository
	Shutdown()
	Transaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type mysqlRepo struct {
	mysql_ports.TransactionRepository
	mysql_ports.ProductRepository

	Store *initializer.Replication
}

func NewMySQL(store *initializer.Replication) SqlInterface {
	sql := &mysqlRepo{
		transaction_mysql_impl.NewTransactionMysqlRepositoryImpl(store),
		product_mysql_impl.NewProductMysqlRepositoryImpl(store),

		store,
	}

	return sql
}

func (repo *mysqlRepo) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return executor.Transaction(ctx, repo.Store, fn)
}

func (repo *mysqlRepo) Shutdown() {
	_ = repo.Store.Primary.Close()
	_ = repo.Store.Standby.Close()
}
