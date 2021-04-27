package main

import (
	"encoding/json"
	"fmt"
	"github.com/xXxRisingTidexXx/wt-labs/pkg/jsonrpc"
)

func main() {
	request := jsonrpc.NewRequest("methods", jsonrpc.NewEmpty(), jsonrpc.NewString("a3sd2d12"))
	bytes, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bytes))
	}
}
