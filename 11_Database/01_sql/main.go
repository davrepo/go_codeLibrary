package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // Add this line to import sqlite3 driver
)

type Product struct {
	Id        int
	Name      string
	Inventory int
	Price     int
}

func main() {
	db, err := sql.Open("sqlite3", "../practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	rows, err := db.Query("SELECT id, name, inventory, price FROM products")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Id, &product.Name, &product.Inventory, &product.Price)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Prodcut: ", product.Id, " ", product.Name, " ", product.Inventory, " ", product.Price)
	}
}
