package misc

import (

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