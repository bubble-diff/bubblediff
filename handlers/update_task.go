package handlers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/bubble-diff/bubblediff/app"
	"github.com/bubble-diff/bubblediff/models"
)

type updateTaskHandler struct {
	Description   string                `json:"description"`
	TrafficConfig *models.TrafficConfig `json:"traffic_config" `
	FilterConfig  *models.FilterConfig  `json:"filter_config" `
	AdvanceConfig *models.AdvanceConfig `json:"advance_config" `
}

var UpdateTaskHandler = &updateTaskHandler{}

func (h *updateTaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	taskid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("parse taskid=%s failed, %s", id, err)
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	err = c.BindJSON(h)
	if err != nil {
		log.Printf("unmarshal json failed, %s", err)
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	err = h.updateTask(c, taskid)
	if err != nil {
		log.Printf("update task failed, %s", err)
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": nil,
	})
}

func (h updateTaskHandler) updateTask(c *gin.Context, taskid int64) (err error) {
	// todo: 对task进行数据检查

	filter := bson.D{{"id", taskid}}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "description", Value: h.Description},
				{Key: "traffic_config", Value: h.TrafficConfig},
				{Key: "filter_config", Value: h.FilterConfig},
				{Key: "advance_config", Value: h.AdvanceConfig},
			},
		},
	}

	result, err := app.TaskColl.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}

	log.Printf("update task ok, taskid: %d, result: %+v", taskid, result)
	return nil
}
