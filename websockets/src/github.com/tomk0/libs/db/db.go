package db

import (


	//"encoding/json"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	misc "github.com/tomk0/libs/misc"

)

type MenuItem struct {

	Id string `json:"Id"`
	Name string `json:"Name"`
	Disc string `json:"Disc"`
	Price float64 `json:"Price"`
	Amount int `json:"Amount"`
	Category string `json:"Category"`


}

/*
func Test(){

	//menu.menu = GetAll()
	//b, err := json.Marshal(menu.menu[0])

	if err != nil{

		panic(err)

	}

	fmt.Println(string(b))


}
*/

func GetAll() []menuItem{


	menu := make([]menuItem, 1)

	db, err := sql.Open("mysql", "tom:pwd123@tcp(127.0.0.1:3306)/cafe_POS_v3")

	misc.CheckError(err)

	defer db.Close()

	results , err := db.Query("SELECT ITM.*, CAT.CAT_TYPE FROM ITEMS AS ITM JOIN ITEM_CATEGORY AS IC ON ITM.ITM_ID = IC.IC_ITM_ID JOIN CATEGORY AS CAT ON IC.IC_CAT_ID = CAT.CAT_ID WHERE ITM.ITM_AMOUNT > 0;")

	misc.CheckError(err)

	for results.Next(){

		var tmp menuItem

		err = results.Scan(&tmp.Id, &tmp.Name, &tmp.Disc, &tmp.Price, &tmp.Amount, &tmp.Category )

		misc.CheckError(err)

		fmt.Println("\n----------------------------------------------------------")
		fmt.Println("ID: ", tmp.Id)
		fmt.Println("Name: ", tmp.Name)
		fmt.Println("Description: ", tmp.Disc)
		fmt.Println("Price: ", tmp.Price)
		fmt.Println("Amount: ", tmp.Amount)
		fmt.Println("Category: ", tmp.Category)

		if menu[0].Id == ""{

			menu[0] = tmp

		}else{

			menu = append(menu, tmp)

		}

	}


	return menu
}

