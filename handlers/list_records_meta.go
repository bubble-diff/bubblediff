package handlers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/bubble-diff/bubblediff/app"
)

func ListRecordsMeta(c *gin.Context) {
	taskid, err := strconv.ParseInt(c.Param("taskid"), 10, 64)
	if err != nil {
		log.Printf("[ListRecordsMeta] parse int failed, %s", err)
		c.JSON(200, gin.H{
			"err":   err.Error(),
			"metas": nil,
		})
		return
	}
	metas, err := app.ListRecordsMeta(c, taskid)
	if err != nil {
		log.Printf("[ListRecordsMeta] get from redis failed, %s", err)
		c.JSON(200, gin.H{
			"err":   err.Error(),
			"metas": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"err":   "",
		"metas": metas,
	})
}
