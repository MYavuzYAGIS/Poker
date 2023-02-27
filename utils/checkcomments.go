package utils

import (
	"fmt"
	"io/ioutil"

	"github.com/fatih/color"
)

func CheckHTMLComments(url string) {
	fmt.Println(" ")
	color.Green("Analyzing Response Body for hidden comments ")
	fmt.Println("===================")

	resp := Connect(url)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("an error occured", err)
	}
	bodyString := string(body)
	fmt.Println(bodyString)
}
