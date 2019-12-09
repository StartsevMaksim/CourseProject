package main

import (
	"database/sql"
	"fmt"
)

func main() {
	database, err := sql.Open("mysql", "root:,jhbcjd@/info")
	if err != nil {
		fmt.Println("Error", err)
	}
	defer database.Close()

	res, err := database.Exec("insert into market.customer(name) value ('Mike')")
	res, err = database.Exec("insert into market.customer(name) value ('Teddy')")
	res, err = database.Exec("insert into market.product(title, price) value ('chair', 120.0)")
	res, err = database.Exec("insert into market.product(title, price) value ('table', 170.5)")
	res, err = database.Exec("insert into market.product(title, price) value ('mirror', 70.0)")
	res, err = database.Exec("insert into market.product(title, price) value ('lamp', 28.6)")
	if err != nil {
		fmt.Println("Error", res, err)
	}
}
