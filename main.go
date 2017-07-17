package main

import "fmt"
import "os"

func main() {
	data := "{\"Hello\":1, \"obj\":[{\"a\":1}]}"

	// TODO get server key from arg
	serverKey := "server key"
	// TODO get registration id from arg
	to := "registration id"
	err := sendData(serverKey, data, to)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
