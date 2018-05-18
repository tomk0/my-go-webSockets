package MyDBStuff

import (


    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
)


/*
type item struct {

    itm_name string `json:"itm_name"`
    itm_price float32 `json:"itm_price"`
    itm_cat string `json:"itm_cat"`

}
*/

func GetDrinks() string{

    db, err := sql.Open("mysql", "tom:pwd123@/Cafe_POS_V2")
    checkErr(err)

    rows, err := db.Query("SELECT I.ITM_NAME, I.ITM_PRICE, C.CAT_NAME FROM ITEMS AS I Join ITEM_CATEGORY AS IC ON I.ITM_ID = IC_ITM_ID JOIN CATEGORY AS C ON IC_CAT_ID = C.CAT_ID WHERE C.CAT_TYPE = 'Drink' AND I.ITM_AMOUNT > 1;")
    checkErr(err)


    var DrinksMenu []item
    for rows.Next(){

        var r structs.item

        err = rows.Scan(&r.itm_name,&r.itm_price, &r.itm_cat)
        checkErr(err)

        fmt.Printf("Item Name: %s, Item Price: %.2f, Item Cat: %s\n", r.itm_name, r.itm_price, r.itm_cat)

        b := json.Marshal(r)

        fmt.Println(string(b))
    }


   // fmt.Printf("Json: %v\n", DrinksMenu)

    return "test"

}

func PassedFrom(Deets string){

    fmt.Println("MyDBStuff.PassedFrom: " + Deets)

}

func checkErr(err error) {
        if err != nil {
            panic(err)
        }
}

