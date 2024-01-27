package main

import "database/sql"

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
