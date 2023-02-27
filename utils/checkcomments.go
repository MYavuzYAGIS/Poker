package utils

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/fatih/color"
)

const COMMENT = "<!--"

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
	matched, _ := regexp.MatchString(COMMENT, bodyString)
	fmt.Println(matched)
}
