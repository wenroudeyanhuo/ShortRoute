// @program:     ShotTener
// @file:        redis.go
// @author:      16574
// @create:      2025-12-15 12:00
// @description:

package sequence

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const sequenceKey = "sequence_stub" // Redis 中的计数器 key，可自定义

type Redis struct {
	client *redis.Redis
}

// NewRedis 创建 Redis 版序列生成器
// nodeConf 为 go-zero 的 redis.RedisConf 配置，例如：
//
//	redis.RedisConf{
//	    Host: "127.0.0.1:6379",
//	    Type: "node",
//	    Pass: "",
//	}
func NewRedis(nodeConf redis.RedisConf) Sequence {
	r := redis.MustNewRedis(nodeConf)
	return &Redis{
		client: r,
	}
}

// Next 获取下一个序列号
// 使用 Redis 的 INCR 命令实现原子自增，与 MySQL 的 REPLACE + LAST_INSERT_ID 效果等价
func (r *Redis) Next() (seq uint64, err error) {
	val, err := r.client.Incr(sequenceKey)
	if err != nil {
		logx.Errorw("redis.Incr failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}
	return uint64(val), nil
}
