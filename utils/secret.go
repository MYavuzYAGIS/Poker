package utils

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

var FILES_TO_CHECK = [17]string{"/robots.txt", "/secret.txt", "/secrets.txt", "/login.html", "/login.php", "/login.aspx", "/cgi-bin", "/img", "/images", "/config.php", "/../../../../../../../etc/passwd", "/api", "/admin", "/cache", "/LICENSE.txt", "/upload", "cron.php"}

func Checkfiles(url string) {
	color.Green("Checking for low hanging urls. \n")
	fmt.Println("=======================")

	for i := 0; i < len(FILES_TO_CHECK); i++ {
		extension := FILES_TO_CHECK[i]
		url = url + extension
		resp := Connect(url)
		statusOK := resp.StatusCode >= 200 && resp.StatusCode < 400

		if !statusOK {
			color.Red("[-] Could not find: %v ", extension)
			continue
		}
		defer resp.Body.Close()

		color.Green("[*] Found: %v", extension)
		color.Green("Reading the Contents of the file : %v", extension)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			color.Red("Could not read the Response Body")
		}
		fmt.Println(string(body))
		fmt.Println("Continuing next targets...")

	}
}
