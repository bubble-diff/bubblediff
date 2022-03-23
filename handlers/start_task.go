package handlers

import (
	"context"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/bubble-diff/bubblediff/app"
)

type startTaskHandler struct{}

var StartTaskHandler = &startTaskHandler{}

func (h *startTaskHandler) StartTask(c *gin.Context) {
	taskid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Printf("[StartTask] parse taskid=%s failed, %s", c.Param("id"), err)
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	err = h.startTask(c, taskid)
	if err != nil {
		log.Printf("[StartTask] start task failed, %s", err)
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
	})
}

func (h *startTaskHandler) startTask(ctx context.Context, taskid int64) (err error) {
	filter := bson.D{{"id", taskid}}
	update := bson.D{
		{Key: "$set", Value: bson.D{{Key: "is_running", Value: true}}},
	}

	result, err := app.TaskColl.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	log.Printf("start task ok, taskid: %d, result: %+v", taskid, result)
	return nil
}
