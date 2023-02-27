package utils

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/fatih/color"
)

var imgRE = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)

func CheckImageNames(url string) {
	fmt.Println(" ")
	color.Green("Checking image names and locations.")
	fmt.Println("===================")

	resp := Connect(url)
	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	if err != nil {
		fmt.Println("an error occured", err)
	}
	imgs := imgRE.FindAllStringSubmatch(bodyString, -1)
	out := make([]string, len(imgs))
	for i := range out {
		out[i] = imgs[i][1]
		color.Green("Image found: %v", out[i])
	}
}
