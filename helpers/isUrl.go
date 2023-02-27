package helpers

import (
	"fmt"
	"net/url"

	"github.com/fatih/color"
)

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	if err != nil {
		fmt.Println("provide scheme")
	}

	if err == nil && u.Scheme != "" && u.Host != "" {
		return true
	}
	color.Red("No Scheme Is Given")
	color.Red("Consider Adding http:// or https:// like http://127.0.0.1")
	return false
}
