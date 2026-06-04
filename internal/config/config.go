// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	//从yaml文件中读取配置，一定要对齐
	ShortUrlDB struct {
		DSN string
	}
	SequenceDB struct {
		DSN string
	}
	BaseString        string   //base62指定基础字符串
	ShortUrlBlackList []string //短链接黑名单  我认为应该做一个语法库来检查
	ShortDomain       string   //短链接的域名

	CacheRedis cache.CacheConf //Redis 缓存
	BloomRedis redis.RedisConf // 布隆过滤器 redis
}
