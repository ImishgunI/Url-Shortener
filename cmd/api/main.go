package main

import (
	"shortener/src/config"
	"shortener/src/storage"
	"shortener/src/url"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Create()
	db_cfg := config.DBConfigCreate()
	db, err := storage.DbConnect(db_cfg)
	if err != nil {
		panic(err)
	}
	h := url.New(db)
	r := gin.Default()
	r.POST("/shorten", h.PostShortUrl)
	err = r.Run(cfg.Port)
	if err != nil {
		panic("Cannot run server on port " + cfg.Port)
	}
	db.Close()
}
