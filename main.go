package main

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

func main() {
	WorkDir := flag.String("w", "./www", "working directory")
	flag.Parse()
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile(*WorkDir, false)))
	r.NoRoute(func(c *gin.Context) {
		str := fmt.Sprintf("%s", c.Request.URL.Query())
		fmt.Println("Queries:", str)
		body, err := ioutil.ReadAll(c.Request.Body)
		if err == nil {
			str2 := fmt.Sprintf("%s", body)
			fmt.Println("Request Body:", str2)

		}
		defaultFile := *WorkDir + "/default"
		if _, err := os.Stat(defaultFile); err == nil {
			//TODO
			c.File(defaultFile)
		} else {
			c.JSON(200, gin.H{"status": "ok"})
		}
	})

	r.Run(":7890")
}
