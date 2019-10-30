package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/tianhai82/stock-timing/rpcs"
)

const domain = "https://stock-timing.web.app"

func main() {
	fmt.Println("Stock Timing starting")
	r := gin.Default()
	domains := []string{domain}
	if gin.Mode() == gin.DebugMode {
		r.Use(static.Serve("/", static.LocalFile("./public", false)))
		domains = append(domains, "http://localhost:8080")
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins: domains,
	}))
	rpcsRouter := r.Group("/rpc")
	rpcs.AddRpcs(rpcsRouter)

	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
