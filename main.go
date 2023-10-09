/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"database/sql"

	"github.com/alanvitalp/go-hexagonal/adapters/db"
	"github.com/alanvitalp/go-hexagonal/application"
	"github.com/alanvitalp/go-hexagonal/cmd"
)

func main() {
	cmd.Execute()

	db2, _ := sql.Open("sqlite3", "./sqlite.db")
	productDbAdapter := db.NewProductDb(db2)
	productService := application.NewProductService(productDbAdapter)
	productService.Create("Product exemplo", 30)
}
