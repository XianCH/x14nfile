package x14nfile

import (
	"github.com/gin-gonic/gin"
)

func HandleDir(c *gin.Context) {
	if FileDir == nil {
		FailWithMsg(c, "file tree is not ready")
		return
	}
	SuccessWithData(c, FileDir)
}
