package main

import (
	"flag"
	"fmt"
)

func main() {
	//---Define the various flags---
	url := flag.String("msg", "Hello, World!",
		"Message to print on console")

	spacingPtr := flag.Bool("spacing", true,
		"Insert a space between messages")
	//---parse the command line into the defined flags---
	flag.Parse()

	fmt.Print(*url)
	if *spacingPtr == true {
		fmt.Print(" ")
	}
	//---returns all the non-flags arguments---
	for _, arg := range flag.Args() {
		fmt.Println(arg, " ")
	}
}
