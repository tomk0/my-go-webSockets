package misc

import (

	"fmt"
	"encoding/json"

	stuc "github.com/tomk0/libs/structs"


)

func CheckError(err error){

	if err != nil {
		panic(err)
	}


}

func JSONCompile (cmd string, Data stuc.DataOut) string {

	CmdOut := stuc.CmdOut{Cmd : cmd, Data : Data}

	jsonEnc, err := json.Marshal(CmdOut)

	CheckError(err)

	return string(jsonEnc)

}

func PrintOutAllOrders(Orders []stuc.OrderOut){


	for _, Order := range Orders{

		fmt.Println("###########################################################################")

		fmt.Println("ID: ", Order.ID)
		fmt.Println("Time: ", Order.Time)
		fmt.Println("Tabel: ", Order.Tabel)
		fmt.Println("Total: ", Order.Total)

		for i, Items := range Order.Items{

			fmt.Println("---------------------------------------------------------------", Order.ID, " ", i)
			fmt.Println("Name: ", Items.Name)
			fmt.Println("Notes: ", Items.Notes)
			fmt.Println("Amount: ", Items.Amount)
		}
	}
}

func PrintOutAnOrder(Order stuc.OrderOut){


	fmt.Println("###########################################################################")

	fmt.Println("ID: ", Order.ID)
	fmt.Println("Time: ", Order.Time)
	fmt.Println("Tabel: ", Order.Tabel)
	fmt.Println("Total: ", Order.Total)

	for i, Items := range Order.Items{

		fmt.Println("---------------------------------------------------------------", Order.ID, " ", i)
		fmt.Println("Name: ", Items.Name)
		fmt.Println("Notes: ", Items.Notes)
		fmt.Println("Amount: ", Items.Amount)
		
	}
}