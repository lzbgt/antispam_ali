package main

import (
	"flag"
	"github.com/gin-gonic/gin"
)

const accessKeyId string = "LTAIlqjD8PZGnpTx"                   // change accesstoken
const accessKeySecret string = "SYjZlAS9FD2a3NoK34cIXoVYGHxkhU" // change accesstoken

var flagConsulAddr string

func init() {
	flag.StringVar(&flagConsulAddr, "consol", "", "consul server address")
}

func main() {
	flag.Parse()
	router := gin.Default()
	router.POST("/:media", HandlePostMedia)
	router.POST("/:media/:rpc", HandlePostMedia)
	router.Run(":8080")
}
