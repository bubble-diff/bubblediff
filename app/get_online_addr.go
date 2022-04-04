package app

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/bubble-diff/bubblediff/config"
)

// GetOnlineAddr 获取diff任务部署的bubblecopy在线上的实例地址
// example: bubblecopy_test_1 -> 127.0.0.1:8888
func GetOnlineAddr(ctx context.Context, taskid int64) (addr string, err error) {
	conf := config.Get()
	key := fmt.Sprintf("bubblecopy_%s_%d", conf.Env, taskid)
	addr, err = rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}
	return addr, err
}
