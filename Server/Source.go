package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type customer struct {
	id   int
	name string
}

type product struct {
	id    int
	title string
	price float64
}

func saveUser(c echo.Context) error {
	database, err := sql.Open("mysql", "root:,jhbcjd@/info")
	if err != nil {
		fmt.Println("Error", err)
	}
	defer database.Close()

	name := c.FormValue("name")
	res, err := database.Exec("insert into market.customer(name) value (?)", name)
	if err != nil {
		fmt.Println("Error", res, err)
	}
	return c.String(http.StatusOK, "ok")
}

func saveProduct(c echo.Context) error {
	database, err := sql.Open("mysql", "root:,jhbcjd@/info")
	if err != nil {
		fmt.Println("Error", err)
	}
	defer database.Close()
	title := c.FormValue("title")
	price := c.FormValue("price")
	res, err := database.Exec("insert into market.product(title, price) value (?, ?)", title, price)
	if err != nil {
		fmt.Println("Error", err, res)
	}
	return c.String(http.StatusOK, "OK")
}

func getProducts(c echo.Context) error {
	array := []product{}
	database, err := sql.Open("mysql", "root:,jhbcjd@/info")
	if err != nil {
		fmt.Println("Error", err)
	}
	res, err := database.Query("select *from market.product")
	for res.Next() {
		var p product
		err = res.Scan(&p.id, &p.price, &p.title)
		if err != nil {
			fmt.Println("Error", err)
		}
		array = append(array, p)
	}
	var str string
	for i := 0; i < len(array); i++ {
		str += strconv.Itoa(array[i].id) + " " + array[i].title + " " + strconv.FormatFloat(array[i].price, 'f', 6, 64) + "\n"
	}
	return c.String(http.StatusOK, str)
}

func getUsers(c echo.Context) error {
	array := []customer{}
	database, err := sql.Open("mysql", "root:,jhbcjd@/info")
	if err != nil {
		fmt.Println("Error", err)
	}
	res, err := database.Query("select *from market.customer")
	for res.Next() {
		var p customer
		err = res.Scan(&p.id, &p.name)
		if err != nil {
			fmt.Println("Error", err)
		}
		array = append(array, p)
	}
	var str string
	for i := 0; i < len(array); i++ {
		str += strconv.Itoa(array[i].id) + " " + array[i].name + "\n"
	}
	return c.String(http.StatusOK, str)
}

func main() {
	e := echo.New()
	e.GET("/products", getProducts)
	e.GET("/users", getUsers)
	e.POST("/addUser", saveUser)
	e.POST("/addProduct", saveProduct)
	e.Logger.Fatal(e.Start(":1321"))
}
