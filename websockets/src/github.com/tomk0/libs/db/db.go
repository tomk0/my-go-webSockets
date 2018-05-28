package db

import (
	"database/sql"
	"fmt"

	misc "github.com/tomk0/libs/misc"
	stuc "github.com/tomk0/libs/structs"

	_ "github.com/go-sql-driver/mysql"
)

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

// GetAllMenu Is the function to Get the whole Menu
func GetAllMenu() []stuc.MenuItemOut {

	menu := make([]stuc.MenuItemOut, 1)

	db, err := sql.Open("mysql", "tom:pwd123@tcp(127.0.0.1:3306)/cafe_POS_v3")

	misc.CheckError(err)

	defer db.Close()

	results, err := db.Query("SELECT ITM.*, CAT.CAT_TYPE FROM ITEMS AS ITM JOIN ITEM_CATEGORY AS IC ON ITM.ITM_ID = IC.IC_ITM_ID JOIN CATEGORY AS CAT ON IC.IC_CAT_ID = CAT.CAT_ID WHERE ITM.ITM_AMOUNT > 0;")

	misc.CheckError(err)

	for results.Next() {

		var tmp stuc.MenuItemOut

		err = results.Scan(&tmp.ID, &tmp.Name, &tmp.Disc, &tmp.Price, &tmp.Amount, &tmp.Category)

		misc.CheckError(err)

		/*
			fmt.Println("\n----------------------------------------------------------")
			fmt.Println("ID: ", tmp.ID)
			fmt.Println("Name: ", tmp.Name)
			fmt.Println("Description: ", tmp.Disc)
			fmt.Println("Price: ", tmp.Price)
			fmt.Println("Amount: ", tmp.Amount)
			fmt.Println("Category: ", tmp.Category)
		*/

		if menu[0].ID == "" {

			menu[0] = tmp

		} else {

			menu = append(menu, tmp)

		}

	}

	fmt.Println("Server: Whole Menu Sent got from database")
	return menu
}

func GetAllOrders() []stuc.OrderOut {

	Orders := make([]stuc.OrderOut, 1)

	db, err := sql.Open("mysql", "tom:pwd123@tcp(127.0.0.1:3306)/cafe_POS_v3")

	misc.CheckError(err)

	defer db.Close()

	results, err := db.Query("SELECT ORD.ORD_ID, ORD.ORD_TIME, ORD.ORD_TOTAL, TBL.TBL_NAME  FROM ORDERS AS ORD JOIN TABLES AS TBL ON ORD.ORD_TBL_ID = TBL.TBL_ID;")

	misc.CheckError(err)

	for results.Next() {

		var tmp stuc.OrderOut

		err = results.Scan(&tmp.ID, &tmp.Time, &tmp.Total, &tmp.Tabel)

		misc.CheckError(err)

			fmt.Println("\n----------------------------------------------------------")
			fmt.Println("ID: ", tmp.ID)
			fmt.Println("Time: ", tmp.Time)
			fmt.Println("Total: ", tmp.Total)
			fmt.Println("Table: ", tmp.Tabel)

		if Orders[0].ID == "" {

			Orders[0] = tmp

		} else {

			Orders = append(Orders, tmp)

		}

	}

	for i, Order := range Orders{

		tmparry := make([]stuc.OrderItemOut, 1)

		fmt.Println(Order.ID)

		results, err = db.Query("SELECT ITM.ITM_NAME, OI.OI_NOTES, OI.OI_AMOUNT FROM ORDER_ITEMS AS OI JOIN ITEMS AS ITM ON OI.OI_ITM_ID = ITM.ITM_ID WHERE OI.OI_ORD_ID = 'ORD00001';")

		misc.CheckError(err)

		for results.Next(){

			var tmpItm stuc.OrderItemOut

			err = results.Scan(&tmpItm.Name, &tmpItm.Notes, &tmpItm.Amount)

			fmt.Println("\n---------------------------------------------------------- ", Order.ID)
			fmt.Println("ID-ITM: ", tmpItm.Name)
			fmt.Println("ID-FILL: ", tmpItm.Notes)
			fmt.Println("Amount: ", tmpItm.Amount)

			if (tmparry[0].ItemID == ""){

				tmparry[0] = tmpItm

			}else{

				tmparry = append(tmparry, tmpItm)
			}

		}

		Orders[i].Items = tmparry

	}

	fmt.Println("Server: All Orders Sent got from database")
	return Orders
}