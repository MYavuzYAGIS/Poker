package utils

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/fatih/color"
)

var KEYWORDS = [20]string{"password", "user", "username", "secret", "private", "$_GET", "hidden", "admin", "login", "pass", "passwd", "hash", "digest", "eval", "innerHTML", "CMD", "include_once", "header", "getenv", "ssh"}

func InterestingKeywors(url string) {
	fmt.Println(" ")
	color.Green("Analyzing for interesting keywords ")
	fmt.Println("===================")
	resp := Connect(url)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("an error occured", err)
	}
	bodyString := string(body)
	for i := 0; i < len(KEYWORDS); i++ {
		matched, _ := regexp.MatchString(KEYWORDS[i], bodyString)
		if !matched {
			color.Red("Not found : %v", KEYWORDS[i])
		} else {
			color.Green("Keyword found : %v", KEYWORDS[i])
		}
	}
}

// TODO: also check urls.
