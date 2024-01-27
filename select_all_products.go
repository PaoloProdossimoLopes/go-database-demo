package main

import "database/sql"

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
