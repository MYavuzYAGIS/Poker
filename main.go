package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/MYavuzYAGIS/Poker/helpers"
	"github.com/fatih/color"
)

func main() {
	// Define the single flag: url

	helpers.Printbanner()

	url := flag.String("url", "Please provide a valid URL WITH http or https", "http://127.0.0.1")
	flag.Parse()
	connect(*url)
}

func connect(url string) {
	helpers.IsUrl(url)
	resp, err := http.Get(url)
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
	color.BlueString("connected to ==> ", url)
	fmt.Println("\n", "====================================")
	fmt.Println(resp)
	fmt.Println("\n", "====================================")
}

// TODO: if url is given without a scheme, try adding scheme(http://) or (https://)
