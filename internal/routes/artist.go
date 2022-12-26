package routes

import (
	"net/http"
	"strconv"

	"github.com/dezhishen/photo-album-metadata/internal/storage/artist"
	"github.com/gin-gonic/gin"
)

func ResigterArtistRoutes(r *gin.Engine) {
	r.GET("/artist/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result, err := artist.Get(int64(id))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	})
}
