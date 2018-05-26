package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

const N = 10

func main() {
	//option 2
	datas2 := make([]Foo, 1)
	for i := 0; i < 10; i++ {
		datas2 = append(datas2, Foo{Number: 1, Title: "test"})
	}
	j, err := json.Marshal(datas2)
	fmt.Println(string(j), err)
	
	fmt.Println(len(datas2))
}
