package executor

import (
	"clean_arch_ws/repository/initializer"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ContextTxKey struct{}

type ContextTxWrapper struct {
	tx *sqlx.Tx
}

func Transaction(ctx context.Context, db *initializer.Replication, fn func(ctx context.Context) error) error {
	// use primary for transaction
	tx, err := db.Primary.Beginx()
	if err != nil {
		return err
	}

	ctxTx := context.WithValue(ctx, ContextTxKey{}, ContextTxWrapper{tx: tx})

	func() {
		defer func() {
			if p := recover(); p != nil {
				// keep original error
				err = tx.Rollback()
				if err != nil {
					fmt.Println(err, "ini rollback error")
				}
				fmt.Println("masok kesiniiiixx", p)
				switch e := p.(type) {
				case error:
					err = e
				default:
					panic(e)
				}
			} else if err != nil {
				fmt.Println(" masok kesini juga")
				_ = tx.Rollback()
			} else {
				fmt.Println("masok commit")
				err = tx.Commit()
			}
		}()
		err = fn(ctxTx)
	}()

	return err
}

// IsTransaction checks wether context contain transaction or not
func IsTransaction(ctx context.Context) (bool, *sqlx.Tx) {
	ctxTxWrapper, ok := ctx.Value(ContextTxKey{}).(ContextTxWrapper)

	return ok, ctxTxWrapper.tx
}
