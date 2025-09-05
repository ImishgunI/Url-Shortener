package url

import (
	"log"
	"log/slog"
	"shortener/src/storage"
	"shortener/src/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Data *storage.DataBase
}

type FactoryRequest interface {
	PostShortUrl(c *gin.Context)
	GetOrigUrl(c *gin.Context)
	PutShortUrl(c *gin.Context)
	DeleteShortUrl(c *gin.Context)
}

type Url struct {
	Url string
}

func New(db *storage.DataBase) *Handler {
	return &Handler{Data: db}
}

func (h *Handler) PostShortUrl(c *gin.Context) {
	var u Url
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
	slog.Info("Url was created succesfully")
	c.JSON(201, gin.H{
		"id":           id,
		"original_url": u.Url,
		"short_url":    short_url,
		"createdAt":    time.Now().Format(time.RFC1123),
		"updetedAt":    time.Now().Format(time.RFC1123),
	})
}

func (h *Handler) GetOrigUrl(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		slog.Error("Url Not Found")
		c.JSON(404, gin.H{"error": "Not Found"})
		return
	}
	id := utils.DecodeBase62(code)
	orig_url, err := h.Data.Get(id)
	if err != nil {
		slog.Error("Orig url not found")
		c.JSON(404, gin.H{"error": "Original URL Not Found"})
		return
	}
	c.JSON(200, gin.H{
		"id":           id,
		"original_url": orig_url,
		"short_url":    code,
		"createdAt":    time.Now().Format(time.RFC1123),
		"updetedAt":    time.Now().Format(time.RFC1123),
	})
}

func (h *Handler) PutShortUrl(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		slog.Error("Short URL Not Found", "URL", code)
		c.JSON(404, gin.H{"error": "Short URL not found"})
		return
	}
	var u Url
	if err := c.BindJSON(&u); err != nil {
		log.Printf("%v", err)
		c.JSON(400, gin.H{"error": "Url not found"})
		return
	}
	id := utils.DecodeBase62(code)
	err := h.Data.Update(u.Url, id)
	if err != nil {
		log.Printf("%v", err)
		c.JSON(500, gin.H{"error": "DataBase wasn't updated"})
		return
	}
	c.JSON(200, gin.H{
		"id":           id,
		"original_url": u.Url,
		"short_url":    code,
		"createdAt":    time.Now().Format(time.RFC1123),
		"updetedAt":    time.Now().Format(time.RFC1123),
	})
}

func (h *Handler) DeleteShortUrl(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		slog.Error("Short URL Not Found", "URL", code)
		c.JSON(404, gin.H{"error": "Short URL not found"})
		return
	}
	id := utils.DecodeBase62(code)
	err := h.Data.Delete(id)
	if err != nil {
		log.Printf("%v", err)
		c.JSON(500, gin.H{"error": "Url wasn't deleted"})
		return
	}
	c.Status(204)
}
