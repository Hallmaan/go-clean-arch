package ucase_transaction

import (
	transaction_domain "clean_arch_ws/internal/entities/transaction"
	ucase_product "clean_arch_ws/internal/usecase/product"
	mysql_ports "clean_arch_ws/repository/mysql/ports"
	nats_ports "clean_arch_ws/repository/nats/ports"
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/nats-io/nats.go"
)

type CreateTransactionUseCasePorts interface {
	Create(ctx context.Context, trx *transaction_domain.TransactionDomain) (*transaction_domain.TransactionDomain, error)
	// create with product dll
}

type AddNewTransactionUCase struct {
	TransactionRepo mysql_ports.TransactionRepository
	GetProductUcase ucase_product.GetProductUcasePorts
	NatsKv          nats_ports.RepositoryNats
	Nats            *nats.Conn
}

func NewAddTrxUsecase(trxRepo mysql_ports.TransactionRepository, pdUcase ucase_product.GetProductUcasePorts, natsKv nats_ports.RepositoryNats, natsJSClient *nats.Conn) CreateTransactionUseCasePorts {
	return &AddNewTransactionUCase{
		TransactionRepo: trxRepo,
		GetProductUcase: pdUcase,
		NatsKv:          natsKv,
		Nats:            natsJSClient,
	}
}

func (trx AddNewTransactionUCase) Create(ctx context.Context, p *transaction_domain.TransactionDomain) (*transaction_domain.TransactionDomain, error) {
	res, err := trx.GetProductUcase.Get(ctx, 1)

	if res == nil || err != nil {
		return nil, errors.New("product not found")
	}

	p.Product = res

	trxCreateId, err := trx.TransactionRepo.Create(ctx, p)

	if err != nil {
		return nil, err
	}

	getTrx, err := trx.TransactionRepo.Get(ctx, trxCreateId)

	if err != nil {
		return nil, err
	}

	// send to nats kv bucket transaction
	str := strconv.FormatInt(getTrx.GetId(), 10)
	err = trx.NatsKv.KVPut("Transaction", str, getTrx)

	if err != nil {
		return nil, err
	}

	message := fmt.Sprintf("Transaction diterima ke-%d", getTrx.GetId())

	err = trx.Nats.Publish("TransactionUpdates", []byte(message))

	if err != nil {
		fmt.Println("error publish :", err)
		return nil, err
	}

	// natsRes, err := trx.NatsKv.KVGet("Transaction", str)

	// if err != nil {
	// 	fmt.Println("error", err)
	// 	return nil, err
	// }

	// xx := &transaction_domain.TransactionDomain{}

	// err = json.Unmarshal(natsRes, xx)

	// if err != nil {
	// 	fmt.Println("nats get", err)
	// 	return nil, err
	// }

	// fmt.Println(xx, "nats get")

	return getTrx, nil
}
