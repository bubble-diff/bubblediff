package handlers

import (
	"context"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/bubble-diff/bubblediff/app"
	"github.com/bubble-diff/bubblediff/models"
)

func GetTaskDetailByID(c *gin.Context) {
	id := c.Param("id")
	taskid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("get task by id=%s failed, caused by parseint: %s", id, err)
		c.JSON(200, gin.H{
			"err":  err.Error(),
			"task": nil,
		})
		return
	}

	task, err := getTaskDetailByID(c, taskid)
	if err != nil {
		log.Printf("get task by id=%d failed, %s", taskid, err)
		c.JSON(200, gin.H{
			"err":  err.Error(),
			"task": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"err":  "",
		"task": task,
	})
}

func getTaskDetailByID(ctx context.Context, taskid int64) (task *models.Task, err error) {
	task = new(models.Task)
	filter := bson.D{{"id", taskid}}
	err = app.TaskColl.FindOne(ctx, filter).Decode(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}
