package db

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/bubble-diff/bubblediff/config"
)

var ctx = context.Background()

var (
	Mongodb *mongo.Client
	Rdb     *redis.Client
)

func Init() (err error) {
	cfg := config.Get()

	ctx1, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	Mongodb, err = mongo.Connect(ctx1, options.Client().ApplyURI(cfg.MongoUrl))
	if err != nil {
		return err
	}

	Rdb = redis.NewClient(&redis.Options{Addr: cfg.Redis.Addr, Password: cfg.Redis.Password})

	log.Println("init db ok.")
	return nil
}
