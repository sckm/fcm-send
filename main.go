package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		to        string
		serverKey string
		data      string
	)

	flag.StringVar(&to, "t", "", "Registration token, Topic name or Device group name")
	flag.StringVar(&serverKey, "s", "", "server key")
	flag.StringVar(&data, "d", "",
		`data payload. ex. {"message":"Hello World!"}`)

	if len(os.Args) == 1 {
		fmt.Println("require arguments\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	if to == "" {
		fmt.Println("require Registration token, Topic name or Device group name.")
		fmt.Println("use -t argument")
		os.Exit(1)
	}

	if serverKey == "" {
		fmt.Println("require server key")
		fmt.Println("use -s argument")
		os.Exit(1)
	}

	if data == "" {
		fmt.Println("require data payload")
		fmt.Println("use -d argument")
		os.Exit(1)
	}

	err := sendData(serverKey, data, to)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
