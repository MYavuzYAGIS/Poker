package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define the single flag: url

	url := flag.String("url", "Please provide a valid URL without http or https", "127.0.0.1")
	flag.Parse()
	fmt.Println("url", *url)

	connect(url)
}

func connect(url *string) {
	resp, err := http.Get(*url)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("connected to ==> ", *url)
	fmt.Println(resp)
}
