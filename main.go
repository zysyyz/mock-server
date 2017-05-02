package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	http "net/http"
	"os"
)

func main() {
	WorkDir := flag.String("w", "./www", "working directory")
	flag.Parse()
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		str := fmt.Sprintf("%s", c.Request.URL.Query())
		fmt.Println("Queries:", str)
		body, err := ioutil.ReadAll(c.Request.Body)
		if err == nil {
			str2 := fmt.Sprintf("%s", body)
			fmt.Println("Request Body:", str2)

		}
		reqFile := *WorkDir + c.Request.URL.Path

		defaultFile := *WorkDir + "/default"
		fmt.Println("Request file" + reqFile)
		content := []byte("{\"status\":\"ok\"}")
		if _, err := os.Stat(reqFile); err == nil {
			content, _ = ioutil.ReadFile(reqFile)
		} else {
			if _, err := os.Stat(defaultFile); err == nil {
				content, _ = ioutil.ReadFile(defaultFile)
			}
		}
		c.Data(http.StatusOK, "application/json", content)
		return
	})

	r.Run(":7890")
}
