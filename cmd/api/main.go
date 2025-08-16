package main

import (
	"shortener/src/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Create()
	r := gin.Default()
	err := r.Run(cfg.Port)
	if err != nil {
		panic("Cannot run server on port " + cfg.Port)
	}
}
