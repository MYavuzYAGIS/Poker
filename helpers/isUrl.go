package helpers

import (
	"net/url"

	"github.com/fatih/color"
)

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	color.Red("No Scheme Is Given")
	color.Red("Consider Adding http:// or https:// like http://127.0.0.1")
	return err == nil && u.Scheme != "" && u.Host != ""
}
