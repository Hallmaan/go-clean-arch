package main

import (
	transaction_controller "clean_arch_ws/pkg/controller/transaction"
	ucase_product "clean_arch_ws/pkg/usecase/product"
	ucase_transaction "clean_arch_ws/pkg/usecase/transaction"
	database_mysql "clean_arch_ws/repository/mysql"
	product_mysql_impl "clean_arch_ws/repository/mysql/impl/product"
	transaction_mysql_impl "clean_arch_ws/repository/mysql/impl/transaction"
	nats_repository_impl "clean_arch_ws/repository/nats/impl"
	transporter "clean_arch_ws/transports/http"
	"clean_arch_ws/transports/http/router"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database_mysql.NewDB()

	/// GET PRODUCT USECASE
	ProductRepo := product_mysql_impl.NewProductMysqlRepositoryImpl(db)
	GetProductUcase := ucase_product.NewGetProductByIdUCase(ProductRepo)
	// ProductController := product_controller.NewGetProductController(GetProductUcase, ProductRepo)

	// res, err := ProductController.Get(1)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(res, "<<<<<< GET PRODUCT USECASE")

	// CREATE TRANSACTION USECASE
	trxRepo := transaction_mysql_impl.NewTransactionMysqlRepositoryImpl(db)
	multipleNatsClient, err := nats_repository_impl.NewMultipleNatsClient()
	if err != nil {
		panic(err)
	}

	natsJSClient, err := nats_repository_impl.NewNatsJetstreamClient()
	if err != nil {
		panic(err)
	}

	natsKv := nats_repository_impl.NewRepositoryNats(natsJSClient, multipleNatsClient)
	createTrxUcase := ucase_transaction.NewAddTrxUsecase(trxRepo, GetProductUcase, natsKv, natsJSClient)
	trxController := transaction_controller.NewCreateTransactionController(createTrxUcase, trxRepo)

	// tx, err := trxController.Create(transaction_domain.TransactionDomain{TrxName: "Transaction 2", Product: res})

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(tx, "<<<<<< CREATE TRANSACTIONS")

	// setupRoutes()

	trxTransporter := transporter.NewTransactionTransporter(trxController)

	router := router.NewRouter(trxTransporter, natsJSClient)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
