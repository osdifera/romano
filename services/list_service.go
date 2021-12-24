package services

import "github.com/gin-gonic/gin"

func GetListFile(c *gin.Context) {
	c.FileAttachment("./file.txt", "file.txt")
}
