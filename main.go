
package main

import (
	"docker-controller/routers"
	"fmt"
	"net/http"
)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr: fmt.Sprintf(":%d", 8080),
		Handler: router,
	}
	s.ListenAndServe()
}