package services

import "github.com/gin-gonic/gin"

func GetListFile(c *gin.Context) {
	c.FileAttachment("./romano.m3u", "romano.m3u")
}
