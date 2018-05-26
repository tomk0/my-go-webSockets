package db

import (


	"encoding/json"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	misc "github.com/tomk0/libs/misc"

)

type menuItem struct {

	id string `json:"id"`
	name string `json:"name"`
	disc string `json:"disc"`
	price float64 `json:"price"`
	amount int `json:"amount"`
	category string `json:"category"`


}


func Test(){

	//menu.menu = GetAll()
	//b, err := json.Marshal(menu.menu[0])

	if err != nil{

		panic(err)

	}

	fmt.Println(string(b))


}

func GetAll() string{


	menu := make([]menuItem, 1)

	db, err := sql.Open("mysql", "tom:pwd123@tcp(127.0.0.1:3306)/cafe_POS_v3")

	misc.CheckError(err)

	defer db.Close()

	results , err := db.Query("SELECT ITM.*, CAT.CAT_TYPE FROM ITEMS AS ITM JOIN ITEM_CATEGORY AS IC ON ITM.ITM_ID = IC.IC_ITM_ID JOIN CATEGORY AS CAT ON IC.IC_CAT_ID = CAT.CAT_ID WHERE ITM.ITM_AMOUNT > 0;")

	misc.CheckError(err)

	for results.Next(){

		var tmp menuItem

		err = results.Scan(&tmp.id, &tmp.name, &tmp.disc, &tmp.price, &tmp.amount, &tmp.category )

		misc.CheckError(err)

		fmt.Println("\n----------------------------------------------------------")
		fmt.Println("ID: ", tmp.id)
		fmt.Println("Name: ", tmp.name)
		fmt.Println("Description: ", tmp.disc)
		fmt.Println("Price: ", tmp.price)
		fmt.Println("Amount: ", tmp.amount)
		fmt.Println("Category: ", tmp.category)

		if menu[0].id == ""{

			menu[0] = tmp

		}else{

			menu = append(menu, tmp)

		}

	}

	j, err := json.Marshel(menu)

	return string(j)
}

