package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"yuequan/routes"
	"yuequan/service"

	"github.com/gin-gonic/gin"
)

func main() {
	go service.OpenXray()
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	r := gin.Default()
	routes.CollectRoute(r)
	r.Run(":9123")
}
