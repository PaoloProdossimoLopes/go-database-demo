package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Driver for MySQL (dont need to use only import here make works)

	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	//db_user:sql_password@tcp(host:port)/db_name
	database, databaseError := sql.Open("mysql", "test:test@tcp(localhost:3306)/test")
	if databaseError != nil {
		panic(databaseError)
	}
	defer database.Close()

	product := NewProduct("Product 1", 10.0)

	insertProductError := insertProduct(database, product)
	if insertProductError != nil {
		panic(insertProductError)
	}

	product.Price = 20.0
	updateProductError := updateProduct(database, product)
	if updateProductError != nil {
		panic(updateProductError)
	}

	findedProduct, selectError := selectOneProduct(database, product.ID)
	if selectError != nil {
		panic(updateProductError)
	}
	fmt.Printf("id=%v\nname=%v\nprice=%v\n", findedProduct.ID, findedProduct.Name, findedProduct.Price)

	products, selectAllError := selectAllProducts(database)
	if selectAllError != nil {
		panic(selectAllError)
	}
	for _, product := range products {
		fmt.Printf("id=%v\nname=%v\nprice=%v\n", product.ID, product.Name, product.Price)
	}

	deleteError := deleteProduct(database, product.ID)
	if deleteError != nil {
		panic(deleteError)
	}
}
