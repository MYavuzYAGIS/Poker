package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func Cookie(url string) {
	color.Green("Checking for cookies.")
	fmt.Println("=======================")
	resp := Connect(url)
	cookies := resp.Cookies()
	if len(cookies) == 0 {
		fmt.Println("No Cookies detected")
	} else {
		fmt.Println("cookies")
		for _, cookie := range resp.Cookies() {
			color.Green("Found a cookie named: %v, %v", cookie.Name, cookie.Value)
		}

	}

	defer resp.Body.Close()
}
