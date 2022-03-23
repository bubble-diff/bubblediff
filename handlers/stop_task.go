package handlers

import (
	"context"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/bubble-diff/bubblediff/app"
)

type stopTaskHandler struct{}

var StopTaskHandler = &stopTaskHandler{}

func (h *stopTaskHandler) StopTask(c *gin.Context) {
	taskid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Printf("[StopTask] parse taskid=%s failed, %s", c.Param("id"), err)
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	err = h.stopTask(c, taskid)
	if err != nil {
		log.Printf("[StopTask] stop task failed, %s", err)
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
	})
}

func (h *stopTaskHandler) stopTask(ctx context.Context, taskid int64) (err error) {
	filter := bson.D{{"id", taskid}}
	update := bson.D{
		{Key: "$set", Value: bson.D{{Key: "is_running", Value: false}}},
	}

	result, err := app.TaskColl.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	log.Printf("stop task ok, taskid: %d, result: %+v", taskid, result)
	return nil
}
