package handlers

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/bubble-diff/bubblediff/app"
	"github.com/bubble-diff/bubblediff/models"
)

func GetTaskDetailByName(c *gin.Context) {
	taskname := c.Query("name")
	task, err := getTaskDetailByName(c, taskname)
	if err != nil {
		log.Printf("get task by %s failed, %s", taskname, err)
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
	}
	c.JSON(200, task)
}

func getTaskDetailByName(ctx context.Context, name string) (task *models.Task, err error) {
	task = new(models.Task)
	filter := bson.D{{"name", name}}
	err = app.TaskColl.FindOne(ctx, filter).Decode(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}
