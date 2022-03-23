package handlers

import (
	"context"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/bubble-diff/bubblediff/app"
)

type deleteTaskHandler struct{}

var DeleteTaskHandler = &deleteTaskHandler{}

func (h *deleteTaskHandler) DeleteTask(c *gin.Context) {
	taskid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Printf("[DeleteTask] parse taskid=%s failed, %s", c.Param("id"), err)
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	err = h.deleteTask(c, taskid)
	if err != nil {
		log.Printf("[DeleteTask] delete task failed, %s", err)
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
	})
}

func (h *deleteTaskHandler) deleteTask(ctx context.Context, taskid int64) (err error) {
	filter := bson.D{{Key: "id", Value: taskid}}
	res, err := app.TaskColl.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	log.Printf("deleted %v documents\n", res.DeletedCount)
	return nil
}
