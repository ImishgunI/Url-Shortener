package url

import (
	"log"
	"shortener/src/storage"
	"shortener/src/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Data *storage.DataBase
}

type ShortenUrl struct {
	Url string
}

func (h *Handler) PostShortUrl(c *gin.Context) {
	var u ShortenUrl
	if err := c.BindJSON(&u); err != nil {
		log.Printf("%v", err)
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	id, err := h.Data.Save(u.Url)
	if err != nil {
		log.Printf("%v", err)
		c.JSON(500, gin.H{"error": "DataBase Not Connected"})
		return
	}

	short_url := utils.EncodeBase62(id)
	log.Println("Url was created succesfully")
	c.JSON(201, gin.H{
		"id":           id,
		"original_url": u.Url,
		"short_url":    short_url,
		"createdAt":    time.Now().Format(time.RFC1123),
		"updetedAt":    time.Now().Format(time.RFC1123),
	})
}

func New(db *storage.DataBase) *Handler {
	return &Handler{Data: db}
}
