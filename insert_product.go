package main

import "database/sql"

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
