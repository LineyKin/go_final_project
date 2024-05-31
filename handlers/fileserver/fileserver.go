package fileserver

import (
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

const webDir = "web"

func FileServerHandler(c *gin.Context) {
	//webDir := config.WebDir()
	filePath := filepath.Join(webDir, strings.TrimPrefix(c.Request.URL.Path, "/"))
	c.File(filePath)
}
