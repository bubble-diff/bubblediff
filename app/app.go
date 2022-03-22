package app

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/bubble-diff/bubblediff/config"
)

var ctx = context.Background()

var (
	UserColl *mongo.Collection
	TaskColl *mongo.Collection
)

var (
	cfg = config.Get()

	mongodb *mongo.Client
	rdb     *redis.Client
)

// IDGenerate 自增id生成器
func IDGenerate(ctx context.Context, namespace string) (int64, error) {
	return rdb.Incr(ctx, fmt.Sprintf("%s_%s_id_gen", cfg.Env, namespace)).Result()
}

func Init() (err error) {
	ctx1, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	mongodb, err = mongo.Connect(ctx1, options.Client().ApplyURI(cfg.Mongo.Url))
	if err != nil {
		return err
	}

	rdb = redis.NewClient(&redis.Options{Addr: cfg.Redis.Addr, Password: cfg.Redis.Password})

	initColl()

	log.Println("init app ok.")
	return nil
}

func initColl() {
	UserColl = mongodb.Database(fmt.Sprintf("bubblediff_%s",cfg.Env)).Collection(cfg.Mongo.Collections.User)
	TaskColl = mongodb.Database(fmt.Sprintf("bubblediff_%s",cfg.Env)).Collection(cfg.Mongo.Collections.Task)
}
