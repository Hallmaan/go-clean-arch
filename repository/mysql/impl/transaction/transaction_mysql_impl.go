package transaction_mysql_impl

import (
	"clean_arch_ws/internal/entities"
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

func (trx TransactionMysqlRepositoryImpl) Get(ctx context.Context, id int64) (*entities.TransactionDomain, error) {
	sql := fmt.Sprintf(`select id, transaction_name from transactions where id = %d`, id)

	trxDomain := &entities.TransactionDomain{}

	err := trx.conn.GetContext(ctx, trxDomain, sql)

	if err != nil {
		return nil, err
	}

	return trxDomain, nil
}

func (trx TransactionMysqlRepositoryImpl) Create(ctx context.Context, transaction *entities.TransactionDomain) (int64, error) {
	sql := `INSERT INTO transactions (transaction_name, product_id) VALUES (?, ?)`

	res, err := trx.conn.ExecContext(ctx, sql, transaction.TrxName, transaction.Product.ID)

	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()

	return id, nil
}
