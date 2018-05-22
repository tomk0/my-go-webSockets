package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type cmdIn struct {
	Cmd string `json:"cmd"`

	Data struct {
		SQL string `json:"SQL"`

		Items []string `json:"items"`
	} `json:"Data"`
}

func jsonPars(msg string) {

	tmp := []byte(msg)
	var r cmdIn

	err := json.Unmarshal(tmp, &r)
	if err != nil {
		panic(err)
	}
	fmt.Println("Raw: ", r)
	fmt.Println("Command: ", r.Cmd)
	fmt.Println("Data: ", r.Data.Items[0])

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

		jsonPars(reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func main() {
	http.Handle("/", websocket.Handler(frmClient))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
