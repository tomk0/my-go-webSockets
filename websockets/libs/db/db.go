package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)

type menuItem struct {

	name string
	cat string
	disc string
	price float64
	amount int


}

var menu []menuItems;// i was here i need to create a gloabl array to put the menu in 

func main(){

	db, err := sql.Open("mysql", "tom:pwd123@tcp(127.0.0.1:3306)/cafe_POS_v3")

	if err != nil {

		panic(err.Error())

	}
	defer db.Close()

	results , err := db.Query("SELECT * FROM `ITEMS` ")

	if err != nil {

		panic(err.Error())

	}



	for results.Next(){

		var uid string
		var name string
		var disc string
		var price float64
		var amount int

		err = results.Scan(&uid, &name, &disc, &price, &amount)

		if err != nil{

			panic(err.Error())

		}

		fmt.Println(uid)

	}
}

func getAllItems(length int){



}
