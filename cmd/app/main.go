package main

import (
	product_controller "clean_arch_ws/pkg/controller/product"
	transaction_controller "clean_arch_ws/pkg/controller/transaction"
	transaction_domain "clean_arch_ws/pkg/domain/transaction"
	ucase_product "clean_arch_ws/pkg/usecase/product"
	ucase_transaction "clean_arch_ws/pkg/usecase/transaction"
	database_mysql "clean_arch_ws/repository/mysql"
	product_mysql_impl "clean_arch_ws/repository/mysql/impl/product"
	transaction_mysql_impl "clean_arch_ws/repository/mysql/impl/transaction"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database_mysql.NewDB()

	/// GET PRODUCT USECASE
	ProductRepo := product_mysql_impl.NewProductMysqlRepositoryImpl(db)
	GetProductUcase := ucase_product.NewGetProductByIdUCase(ProductRepo)
	ProductController := product_controller.NewGetProductController(GetProductUcase, ProductRepo)

	res, err := ProductController.Get(1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res, "<<<<<< GET PRODUCT USECASE")

	// GET TRANSACTION USECASE
	trxRepo := transaction_mysql_impl.NewTransactionMysqlRepositoryImpl(db)
	createTrxUcase := ucase_transaction.NewAddTrxUsecase(trxRepo, GetProductUcase)
	trxController := transaction_controller.NewCreateTransactionController(createTrxUcase, trxRepo)

	tx, err := trxController.Create(transaction_domain.TransactionDomain{TrxName: "Transaction 2", Product: res})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tx, "<<<<<< CREATE TRANSACTIONS")

}
