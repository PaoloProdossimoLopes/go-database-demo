package main

import "database/sql"

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
