package main

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	s := &http.Server{
		Addr:			":8080",
		Handler: 		r,
		ReadTimeout: 	10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}