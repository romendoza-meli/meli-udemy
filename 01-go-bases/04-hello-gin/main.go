package main

import (
	"github.com/gin-gonic/gin"
)
func SayHello(ctx *gin.Context) {
	ctx.String(200, "Hello world!")
}

func SayHelloWithName(ctx *gin.Context) {
	ctx.String(200, "Hello " + ctx.Param("name"))
}

func main (){
	//Server creation with gin Default
	server := gin.Default()

	server.GET("/hello", SayHello)
	server.GET("/hello/:name", SayHelloWithName)
	server.Run(":8080")

}