package transaction_mysql_impl

import (
	"clean_arch_ws/internal/entities"
	"clean_arch_ws/repository/initializer"
	"clean_arch_ws/repository/mysql/executor"
	mysql_ports "clean_arch_ws/repository/mysql/ports"
	"context"
	"fmt"
)

type TransactionMysqlRepositoryImpl struct {
	conn *initializer.Replication
}

func NewTransactionMysqlRepositoryImpl(conn *initializer.Replication) mysql_ports.TransactionRepository {
	return &TransactionMysqlRepositoryImpl{
		conn: conn,
	}
}

func (trx TransactionMysqlRepositoryImpl) GetTransaction(ctx context.Context, id int64) (*entities.TransactionDomain, error) {
	sql := fmt.Sprintf(`select id, transaction_name from transactions where id = %d`, id)

	trxDomain := &entities.TransactionDomain{}

	err := trx.conn.Standby.GetContext(ctx, trxDomain, sql)

	if err != nil {
		return nil, err
	}

	return trxDomain, nil
}

func (trx TransactionMysqlRepositoryImpl) CreateTransaction(ctx context.Context, transaction *entities.TransactionDomain) (int64, error) {
	sql := `INSERT INTO transactions (transaction_name, product_id) VALUES (?, ?)`

	ok, tx := executor.IsTransaction(ctx)

	fmt.Println("ok, tx", ok, tx)

	if ok {
		resTrx, err := tx.ExecContext(ctx, sql, transaction.TrxName, transaction.Product.ID)
		if err != nil {
			fmt.Println(err)
		}

		idTrx, _ := resTrx.LastInsertId()

		return idTrx, nil
	}

	res, err := trx.conn.Primary.ExecContext(ctx, sql, transaction.TrxName, transaction.Product.ID)

	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()

	return id, nil
}
