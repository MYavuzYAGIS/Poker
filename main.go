package main

import (
	"flag"
	"fmt"
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
		color.Red("Cannot connect to the given URL")
		return
	}
	defer resp.Body.Close()
	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 400
	if !statusOK {
		fmt.Println(resp.StatusCode, "returned as status Code. Connection failed")
		return
	}
	color.BlueString("connected to ==> ", url)
	fmt.Println("\n", "====================================")
}
