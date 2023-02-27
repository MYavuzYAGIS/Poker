package utils

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/fatih/color"
)

var imgRE = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)

var inline = regexp.MustCompile((`<script.*?>(.*)</script>`))

var JS = regexp.MustCompile(`<script[^>]+\bsrc=["']([^"']+)["']`)

func CheckImageNames(url string) {
	fmt.Println(" ")
	color.Blue("Checking image names and locations.")
	fmt.Println("===================")

	resp := Connect(url)
	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	if err != nil {
		fmt.Println("an error occured", err)
	}
	imgs := imgRE.FindAllStringSubmatch(bodyString, -1)
	scripts := JS.FindAllStringSubmatch(bodyString, -1)
	inline := inline.FindAllStringSubmatch(bodyString, -1)
	out := make([]string, len(imgs))
	js := make([]string, len(scripts))
	injs := make([]string, len(inline))

	for i := range out {
		out[i] = imgs[i][1]
		color.Green("Image found: %v", out[i])
	}
	color.Blue("Checking for JavaScript files")
	fmt.Println("===================")

	for i := range js {
		js[i] = scripts[i][1]
		color.Green("Javascript Files found: %v", js[i])
	}

	color.Blue("Checking for inline Javascript")
	fmt.Println("===================")

	for i := range injs {
		injs[i] = inline[i][1]
		color.Green("Inline JS found: %v", injs[i])
	}
}
