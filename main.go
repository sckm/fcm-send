package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var (
		to        string
		serverKey string
		data      string
		dataPath  string
	)

	flag.StringVar(&to, "t", "", "Registration token, Topic name or Device group name")
	flag.StringVar(&serverKey, "s", "", "server key")
	flag.StringVar(&data, "d", "",
		`data payload. ex. {"message":"Hello World!"}`)
	flag.StringVar(&dataPath, "p", "",
		fmt.Sprintf("json file for data payload. file content example is: %s", `{"message":"Hello"}`))

	if len(os.Args) == 1 {
		fmt.Println("require arguments")
		fmt.Println()
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

	if dataPath != "" {
		d, err := ioutil.ReadFile(dataPath)
		if err != nil {
			fmt.Printf("failed load data from %s\n", dataPath)
			os.Exit(1)
		}
		data = string(d)
	} else {
		if data == "" {
			fmt.Println("require data payload")
			fmt.Println("use -d or -p argument")
			os.Exit(1)
		}
	}

	err := sendData(serverKey, data, to)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
