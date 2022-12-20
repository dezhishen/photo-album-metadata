package routes

import (
	"net/http"

	"github.com/dezhishen/photo-album-metadata/pkg/model"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) error {
	r.GET("/artist", func(c *gin.Context) {
		artist := model.Artist{}
		c.JSON(http.StatusOK, artist)
	})
	return nil
}
