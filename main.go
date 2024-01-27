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

func insertProduct(db *sql.DB, product *Product) error {
	//`Prepare` method voiding sql injection
	stmt, stmtError := db.Prepare("insert into products(id, name, price) values(?,?,?)")
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

func updateProduct(db *sql.DB, product *Product) error {
	//`Prepare` method voiding sql injection
	stmt, stmtError := db.Prepare("update products set name = ?, price = ? where id = ?")
	if stmtError != nil {
		return stmtError
	}
	defer stmt.Close()

	//replace variables in stmt for values
	_, insertError := stmt.Exec(product.Name, product.Price, product.ID)
	if insertError != nil {
		return insertError
	}

	return nil
}

func selectOneProduct(db *sql.DB, id string) (*Product, error) {
	//`Prepare` method voiding sql injection
	stmt, stmtError := db.Prepare("select id, name, price from products where id = ?")
	if stmtError != nil {
		return nil, stmtError
	}
	defer stmt.Close()

	var product Product
	selectError := stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if selectError != nil {
		return nil, selectError
	}

	return &product, nil
}

func selectAllProducts(db *sql.DB) ([]*Product, error) {
	// Como nao tem injeçao de valories na query do banco de dados nao
	// tem a nescessidade de previnir "SQL Injection" com o uso do metodo "Prepare"
	// podendo assim fazer a chama direta
	rows, rowsError := db.Query("select id, name, price from products")
	if rowsError != nil {
		return nil, rowsError
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		var product Product
		scanError := rows.Scan(&product.ID, &product.Name, &product.Price)
		if scanError != nil {
			return nil, scanError
		}
		products = append(products, &product)
	}

	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	//`Prepare` method voiding sql injection
	stmt, stmtError := db.Prepare("delete from products where id = ?")
	if stmtError != nil {
		return stmtError
	}
	defer stmt.Close()

	//replace variables in stmt for values
	_, insertError := stmt.Exec(id)
	if insertError != nil {
		return insertError
	}

	return nil
}
