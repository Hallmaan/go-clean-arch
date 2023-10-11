package main

import (
	transaction_controller "clean_arch_ws/internal/controller/transaction"
	ucase_product "clean_arch_ws/internal/usecase/product"
	ucase_transaction "clean_arch_ws/internal/usecase/transaction"
	"clean_arch_ws/repository/initializer"
	mysqlrepo "clean_arch_ws/repository/mysql"
	nats_repository_impl "clean_arch_ws/repository/nats/impl"
	transporter "clean_arch_ws/transports/http"
	"clean_arch_ws/transports/http/router"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := initializer.NewMySQLInit()
	mysqlRepo := mysqlrepo.NewMySQL(db)

	/// GET PRODUCT USECASE
	GetProductUcase := ucase_product.NewGetProductByIdUCase(mysqlRepo)

	// CREATE TRANSACTION USECASE
	multipleNatsClient, err := nats_repository_impl.NewMultipleNatsClient()
	if err != nil {
		panic(err)
	}

	natsJSClient, err := nats_repository_impl.NewNatsJetstreamClient()
	if err != nil {
		panic(err)
	}

	natsKv := nats_repository_impl.NewRepositoryNats(natsJSClient, multipleNatsClient)
	createTrxUcase := ucase_transaction.NewAddTrxUsecase(mysqlRepo, GetProductUcase, natsKv, multipleNatsClient)
	trxController := transaction_controller.NewCreateTransactionController(createTrxUcase)
	trxTransporter := transporter.NewTransactionTransporter(trxController)

	router := router.NewRouter(trxTransporter, multipleNatsClient)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
