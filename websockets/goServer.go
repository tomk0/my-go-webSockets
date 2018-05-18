package main

import (

    "./libs/structs"
    "./libs/db"
    "golang.org/x/net/websocket"
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)

type rawJson struct {

    Command string
    Details string

}

func jsonDecode(enc string) rawJson{

    var raw rawJson
    json.Unmarshal([]byte(enc), &raw)
    fmt.Printf("Json Decoder = Command: %s, Details: %s\n", raw.Command, raw.Details)

    return raw
}

func router(ws *websocket.Conn) {
    var err error
    var msg string
    var decode rawJson

    for {
        var reply string

// getting message from client
        if err = websocket.Message.Receive(ws, &reply); err != nil {
            fmt.Println("Can't receive")
            break
        }

        //function to decode the json reply
        decode = jsonDecode(reply)

        //debug to print command
        fmt.Println("Received back from client: " + decode.Command)

        //handeling the command
        switch  decode.Command {
        case "GetDrinks":
            MyDBStuff.GetDrinks()
        case "ping":
            msg = "pong"
        case "Test":
            msg = "Test2You"
        default :
            msg = "Echo: "+ reply

        }

        fmt.Println("Sending to client: " + msg)

        if err = websocket.Message.Send(ws, msg); err != nil {
            fmt.Println("Can't send")
            break
        }
    }
}

func main() {

    MyDBStuff.GetDrinks()

    //handeling the websocket
    http.Handle("/", websocket.Handler(router))
    //erro handeling
    if err := http.ListenAndServe(":1234", nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
