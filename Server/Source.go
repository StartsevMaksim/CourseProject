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
	id int
	name string
}

type product struct {
	id int
	title string
	price float64
}

func save (c echo.Context) error{
	database, err:=sql.Open("mysql", "root:,jhbcjd@/info")
	name:=c.FormValue("name")
	if err!=nil{
		fmt.Println("Error", err)
	}
	defer database.Close()
	res, err:=database.Exec("insert into market.customer(name) value (?)", name)
	if err!=nil{
		fmt.Println("Error", res, err)
	}
	return c.String(http.StatusOK, "ok")
}

func getUser (c echo.Context) error{
	array:=[]product{}
	database, err:=sql.Open("mysql", "root:,jhbcjd@/info")
	if err!=nil{
		fmt.Println("Error", err)
	}
	res, err:=database.Query("select *from market.product")
	for res.Next(){
		var p product
		err=res.Scan(&p.id, &p.price, &p.title)
		if err!=nil{
			fmt.Println("Error", err)
		}
		array=append(array, p)
	}
	var str string
	for i:=0; i<4; i++{
		str+=strconv.Itoa(array[i].id)+" "+array[i].title+" "+strconv.FormatFloat(array[i].price, 'f', 6, 64)+"\n"
	}
	return c.String(http.StatusOK, str)
}

func main(){
	database, err:=sql.Open("mysql", "root:,jhbcjd@/info")
	if err!=nil{
		fmt.Println("Error", err)
	}
	defer database.Close()

	res, err:=database.Exec("insert into market.product(title, price) value ('chair', 120.0)")
	res, err=database.Exec("insert into market.product(title, price) value ('table', 170.5)")
	res, err=database.Exec("insert into market.product(title, price) value ('mirror', 70.0)")
	res, err=database.Exec("insert into market.product(title, price) value ('lamp', 28.6)")
	if err!=nil{
		fmt.Println("Error", res, err)
	}

	e:=echo.New()
	e.GET("/products", getUser)
	e.Logger.Fatal(e.Start(":1321"))
}


