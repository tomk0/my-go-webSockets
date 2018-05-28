package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	DB "github.com/tomk0/libs/db"
	misc "github.com/tomk0/libs/misc"
	stuc "github.com/tomk0/libs/structs"
	"golang.org/x/net/websocket"
)

func menu() string {

	menu := DB.GetAllMenu()
	MenuOut := stuc.MenuOut{ItemsAry: menu}
	Data := stuc.DataOut{Used: true, Menu: MenuOut}

	return misc.JSONCompile("FullMenu", Data)

}

func route(data stuc.CmdIn) string {

	Data := stuc.DataOut{Used: false}

	switch data.Cmd {
	case "ping":
		DB.GetAllOrders()
		return pong(data.Cmd)
	case "Ping":
		return pong(data.Cmd)
	case "getMenu":
		return menu()
	default:
		return misc.JSONCompile("Not a Valid input", Data)

	}

}

func pong(cmd string) string {

	Data := stuc.DataOut{Used: false}

	switch cmd {
	case "ping":
		return misc.JSONCompile("pong", Data)
	case "Ping":
		return misc.JSONCompile("pong", Data)
	default:
		return misc.JSONCompile("", Data)

	}

}

func jsonPars(msg string) stuc.CmdIn {

	tmp := []byte(msg)
	var r stuc.CmdIn

	err := json.Unmarshal(tmp, &r)

	misc.CheckError(err)

	return r

}

func frmClient(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		data := jsonPars(reply)

		resp := route(data)

		if resp == "" {

		} else {

			msg := resp
			//fmt.Println("Sending to client: " + msg)

			if err = websocket.Message.Send(ws, msg); err != nil {
				fmt.Println("Can't send")
				break
			}
		}
	}
}

func main() {

	http.Handle("/", websocket.Handler(frmClient))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
