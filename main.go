package main

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"io/ioutil"
)

func main() {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		str := fmt.Sprintf("%s", c.Request.URL.Query())
		fmt.Println("Queries:", str)
		body, err := ioutil.ReadAll(c.Request.Body)
		if err == nil {
			str2 := fmt.Sprintf("%s", body)
			fmt.Println("Request Body:", str2)

		}
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.Run(":7890")
}
