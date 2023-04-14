package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uchupx/bpjs-test-golang/config"
	"github.com/uchupx/bpjs-test-golang/transport"
)

func main() {
	conf := config.GetConfig()
	trans := transport.Transport{}
	transactionHandler := trans.GetTransactionHandler(conf)
	router := gin.Default()

	router.POST("/transactions", transactionHandler.Posts)
	router.GET("/transactions", transactionHandler.Get)
	router.Run(":8081")
}
