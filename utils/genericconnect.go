package utils

import (
	"fmt"
	"net/http"
)

func Connect(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	return resp
}
