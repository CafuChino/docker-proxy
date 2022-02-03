package main

import (
	"docker-controller/conf"
	"docker-controller/routers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	conf.LoadConfig("");
	config := conf.Conf;
	if config.Server.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := routers.InitRouter()
	s := &http.Server{
		Addr: fmt.Sprintf("%s:%d",config.Server.Address, config.Server.Port),
		Handler: router,
	}
	s.ListenAndServe()
}