package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver for MySQL (dont need use only import works)

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
	//db_user:db_password@tcp(host:port)/db_name
	database, databaseError := sql.Open("mysql", "test:root@tcp(localhost:3306)/test")
	if databaseError != nil {
		panic(databaseError)
	}
	defer database.Close()

	insertProductError := insertProduct(database, NewProduct("Product 1", 10.0))
	if insertProductError != nil {
		panic(insertProductError)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, stmtError := db.Prepare("insert into products(id, name, price) values(?,?,?)") //voiding sql injection
	if stmtError != nil {
		return stmtError
	}
	defer stmt.Close()

	//replace variables in stmt for values
	_, insertError := stmt.Exec(product.ID, product.Name, product.Price)
	if insertError != nil {
		return insertError
	}

	return nil
}
