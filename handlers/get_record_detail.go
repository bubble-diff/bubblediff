package handlers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/bubble-diff/bubblediff/app"
)

func GetRecordDetail(c *gin.Context) {
	taskID, err := strconv.ParseInt(c.Param("taskid"), 10, 64)
	if err != nil {
		log.Printf("[GetRecordDetail] parse int failed, %s", err)
		c.JSON(200, gin.H{
			"err":    err.Error(),
			"record": nil,
		})
		return
	}
	recordID := c.Param("recordid")
	record, err := app.DownloadRecord(c, taskID, recordID)
	if err != nil {
		log.Printf("[GetRecordDetail] download record from cos failed, %s", err)
		c.JSON(200, gin.H{
			"err":    err.Error(),
			"record": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"err": nil,
		"record": gin.H{
			"task_id":   record.TaskID,
			"old_req":   string(record.OldReq),
			"old_resp":  string(record.OldResp),
			"new_resp":  string(record.NewResp),
			"diff":      record.Diff,
			"diff_rate": fmt.Sprintf("%.2f%%", record.DiffRate),
		},
	})
}
