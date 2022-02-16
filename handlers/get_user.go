package handlers

import (
	"context"
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/bubble-diff/bubblediff/db"
	"github.com/bubble-diff/bubblediff/models"
)

func GetUser(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	uid := int(claims[jwt.IdentityKey].(float64))

	user, err := getUserFromMongodb(uid)
	if err != nil {
		log.Printf("get user failed, %s", err)
		c.JSON(400, gin.H{
			"msg": "no such user.",
		})
		return
	}

	c.JSON(200, user)
}

func getUserFromMongodb(uid int) (user *models.User, err error) {
	user = new(models.User)
	coll := db.Mongodb.Database("bubblediff_test").Collection("user")
	err = coll.FindOne(context.Background(), bson.D{{"id", uid}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
