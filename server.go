package x14nfile

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Serverstart() {
	r := gin.Default()
	r.GET("/dir", HandleDir)
	r.Run(":8080")
	fmt.Println("server started")

}
