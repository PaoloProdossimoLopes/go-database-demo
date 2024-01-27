package main

import "database/sql"

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
