package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/bubble-diff/bubblediff/app"
	"github.com/bubble-diff/bubblediff/models"
)

// taskPreview 任务概览
type taskPreview struct {
	ID       int64        `json:"id" bson:"id"`
	Name     string       `json:"name" bson:"name"`
	Owner    *models.User `json:"owner" bson:"owner"`
	TotalReq int          `json:"total_req"`
	SucRate  float64      `json:"suc_rate"`
	DiffRate float64      `json:"diff_rate"`
}

// GetTasks 返回所有diff任务信息概要。
// 支持搜索条件：任务owner，任务名称
func GetTasks(c *gin.Context) {
	owner := c.Query("owner")
	search := c.Query("search")
	tasks, err := getTasks(c, owner, search)
	if err != nil {
		log.Printf("get tasks failed, %s", err)
		c.JSON(200, gin.H{
			"err":         err.Error(),
			"tasks":       nil,
			"total_count": 0,
		})
		return
	}
	c.JSON(200, gin.H{
		"err":         "",
		"tasks":       tasks,
		"total_count": len(tasks),
	})
}

// getTasks 返回符合条件的所有任务，忽略空的owner/search。
func getTasks(ctx context.Context, owner, search string) (tasks []taskPreview, err error) {
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"id", 1}})

	if len(owner) != 0 {
		filter = append(filter, bson.E{Key: "owner.login", Value: owner})
	}

	// 模糊查询
	if len(search) != 0 {
		filter = append(filter, bson.E{
			Key: "name",
			Value: bson.D{
				bson.E{
					Key: "$regex",
					Value: primitive.Regex{
						Pattern: fmt.Sprintf(".*%s.*", search),
						Options: "i", // Case insensitivity
					},
				},
			},
		})
	}

	cur, err := app.TaskColl.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var preview taskPreview
		err = cur.Decode(&preview)
		if err != nil {
			return nil, err
		}
		// todo: 从redis获取请求meta

		tasks = append(tasks, preview)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
