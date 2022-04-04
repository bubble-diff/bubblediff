package handlers

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/bubble-diff/bubblediff/app"
	"github.com/bubble-diff/bubblediff/models"
)

const invalidID int64 = -1

func AddTask(c *gin.Context) {
	task := new(models.Task)
	err := c.BindJSON(task)
	if err != nil {
		log.Printf("unmarshal json failed, %s", err)
		c.JSON(200, gin.H{
			"err": err.Error(),
			"id":  invalidID,
		})
		return
	}

	id, err := addTask(c, task)
	if err != nil {
		log.Printf("add task failed, %s", err)
		c.JSON(200, gin.H{
			"err": err.Error(),
			"id":  invalidID,
		})
		return
	}

	c.JSON(200, gin.H{
		"err": nil,
		"id":  id,
	})
}

func addTask(c *gin.Context, task *models.Task) (id int64, err error) {
	// todo: 对task进行数据检查

	id, err = app.IDGenerate(c, "task")
	if err != nil {
		return invalidID, err
	}

	task.ID = id
	task.CreatedTime = time.Now().Format("2006-01-02 15:04:05")
	task.UpdatedTime = task.CreatedTime

	res, err := app.TaskColl.InsertOne(c, task)
	if err != nil {
		return invalidID, err
	}
	log.Printf("add task ok, object_id: %d, id: %d", res.InsertedID, id)
	return id, nil
}
