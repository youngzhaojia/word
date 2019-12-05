package main

import (
	"fmt"
	"net/http"
	"word/pkg/setting"
	"word/routers"
)

func main() {
	router := routers.InitRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
