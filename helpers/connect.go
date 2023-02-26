package helpers

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

func Checkconnect(url string) {
	IsUrl(url)
	resp, err := http.Get(url)
	if err != nil {
		color.Red("Cannot connect to the given URL")
		return
	}
	defer resp.Body.Close()
	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 400
	if !statusOK {
		fmt.Println(resp.StatusCode, "returned as status Code. Connection failed")
		return
	}
	fmt.Println("connected to ==> ", url)
	fmt.Println("\n", "====================================")
}
