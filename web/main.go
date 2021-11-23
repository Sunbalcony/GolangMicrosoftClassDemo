package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func index(ctx *gin.Context)  {
	fmt.Println(ctx.Request.Body)


}
func main()  {
	r:=gin.Default()
	r.POST("/",index)
	r.Run("0.0.0.0:8080")

}
