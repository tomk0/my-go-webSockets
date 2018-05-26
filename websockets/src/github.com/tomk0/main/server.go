package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	misc "github.com/tomk0/libs/misc"
	DB "github.com/tomk0/libs/db"
	"golang.org/x/net/websocket"
)

type Data struct {

	SQL      string  `json:"SQL"`
	ItemsAry []DB.MenuItem `json:"Items"`

}

type cmdIn struct {
	Cmd  string `json:"Cmd"`
	Data Data `json:"Data"`
}

func route(data cmdIn) string {

	switch data.Cmd {
	case "ping":
		return pong(data.Cmd)
	case "Ping":
		return pong(data.Cmd)
	default:
		return "Not a Valid input"

	}

}

func pong(cmd string) string {

	switch cmd {
	case "ping":
		return "pong"
	case "Ping":
		return "Pong"
	default:
		return ""

	}

}

func jsonPars(msg string) cmdIn {

	tmp := []byte(msg)
	var r cmdIn

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
			fmt.Println("Sending to client: " + msg)

			if err = websocket.Message.Send(ws, msg); err != nil {
				fmt.Println("Can't send")
				break
			}
		}
	}
}

func main() {

	 test := DB.GetAll()
	 Data := Data{SQL: "", ItemsAry : test}
	 test2 := cmdIn{Cmd: "Test", Data: Data}

	 test3, err := json.Marshal(test2) 

	 misc.CheckError(err)

	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n", string(test3))

	http.Handle("/", websocket.Handler(frmClient))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
