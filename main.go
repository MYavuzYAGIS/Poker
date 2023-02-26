package main

import (
	"flag"

	"github.com/MYavuzYAGIS/Poker/helpers"
	"github.com/MYavuzYAGIS/Poker/utils"
)

func main() {
	// Define the single flag: url

	helpers.Printbanner()

	url := flag.String("url", "Please provide a valid URL WITH http or https", "http://127.0.0.1")
	flag.Parse()
	helpers.Checkconnect(*url)
	utils.Checkfiles(*url)
	utils.Cookie(*url)
}
