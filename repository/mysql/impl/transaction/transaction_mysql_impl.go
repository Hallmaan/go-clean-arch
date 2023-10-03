package transaction_mysql_impl

import (
	transaction_domain "clean_arch_ws/pkg/domain/transaction"
	mysql_ports "clean_arch_ws/repository/mysql/ports"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TransactionMysqlRepositoryImpl struct {
	conn *sqlx.DB
}

func NewTransactionMysqlRepositoryImpl(conn *sqlx.DB) mysql_ports.TransactionRepository {
	return &TransactionMysqlRepositoryImpl{
		conn: conn,
	}
}

func (trx TransactionMysqlRepositoryImpl) Get(ctx context.Context, id int64) (*transaction_domain.TransactionDomain, error) {
	sql := fmt.Sprintf(`select id, transaction_name from transactions where id = %d`, id)
	fmt.Println(sql, "ini sqlnya")

	trxDomain := &transaction_domain.TransactionDomain{}

	err := trx.conn.GetContext(ctx, trxDomain, sql)

	if err != nil {
		return nil, err
	}

	return trxDomain, nil
}

func (trx TransactionMysqlRepositoryImpl) Create(ctx context.Context, transaction *transaction_domain.TransactionDomain) (int64, error) {
	sql := `INSERT INTO transactions (transaction_name, product_id) VALUES (?, ?)`

	res, err := trx.conn.ExecContext(ctx, sql, transaction.TrxName, transaction.Product.ID)

	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()

	return id, nil
}
