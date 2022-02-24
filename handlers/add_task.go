package handlers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/bubble-diff/bubblediff/app"
	"github.com/bubble-diff/bubblediff/models"
)

const invalidID int64 = -1

func AddTask(c *gin.Context) {
	task := new(models.Task)
	err := c.BindJSON(task)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": fmt.Sprintf("unmarshal json failed, %s", err),
		})
	}

	id, err := addTask(c, task)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": fmt.Sprintf("add task failed, %s", err),
		})
	}

	c.JSON(200, gin.H{
		"id": id,
	})
}

func addTask(c *gin.Context, task *models.Task) (id int64, err error) {
	id, err = app.IDGenerate(c, "task")
	if err != nil {
		return invalidID, err
	}
	task.ID = id
	res, err := app.TaskColl.InsertOne(c, task)
	if err != nil {
		return invalidID, err
	}
	log.Printf("add task ok, object_id: %d, id: %d", res.InsertedID, id)
	return id, nil
}
