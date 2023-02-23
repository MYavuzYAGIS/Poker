package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
)

func main() {
	// Define the single flag: url

	url := flag.String("url", "Please provide a valid URL without http or https", "127.0.0.1")
	flag.Parse()
	connect(url)
}

func connect(url *string) {
	resp, err := http.Get(*url)
	if err != nil {
		log.Fatalln(err)
		color.Red("Cannot connect to the given URL")
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 400
	if !statusOK {
		fmt.Println("Non-OK HTTP status:", resp.StatusCode)
		return
	}
	color.Blue("connected to ==> ", *url)
	fmt.Println("\n", "====================================")
	fmt.Println(resp)
	fmt.Println("\n", "====================================")
}

// TODO: if url is given without a scheme, try adding scheme(http://) or (https://)
