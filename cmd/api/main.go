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
	r.GET("/shorten/:code", h.GetOrigUrl)
	r.PUT("/shorten/:code", h.PutShortUrl)
	r.DELETE("/shorten/:code", h.DeleteShortUrl)
	err = r.Run(cfg.Port)
	if err != nil {
		panic("Cannot run server on port " + cfg.Port)
	}
	db.Close()
}
