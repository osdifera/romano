package lists

import (
	"github.com/gin-gonic/gin"
	"github.com/osdifera/romano/services"
)

func GetList(c *gin.Context) {
	services.GetListFile(c)
}
