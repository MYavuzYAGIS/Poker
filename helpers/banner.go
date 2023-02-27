package helpers

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func Printbanner() {
	fmt.Println("\t o--o      o  ")
	fmt.Println("\t |   |     | / ")
	fmt.Println("\t O--o  o-o OO   o-o o-o ")
	fmt.Println("\t |     | | | \\  |-' |")
	fmt.Println("\t o     o-o o  o o-o o ")
	color.Red("Automates trivial enumeration \t https://github.com/MYavuzYAGIS")
	time.Sleep(1 * time.Second)
}
