// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"ShotTener/internal/config"
	"ShotTener/model"
	"ShotTener/sequence"
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// 注意一定要大写
type ServiceContext struct {
	Config         config.Config
	ShortUrlModel  model.ShortUrlMapModel
	Sequence       sequence.Sequence
	ShortBlackList map[string]struct{}
	//Bloom filter
	Filter *bloom.Filter
}

func NewServiceContext(c config.Config) *ServiceContext {
	//先连接上数据库才能给下游提供 mysql的服务
	conn := sqlx.NewMysql(c.ShortUrlDB.DSN)
	//把配置文件中黑名单加载到map中,方便后续判断
	m := make(map[string]struct{}, len(c.ShortUrlBlackList))
	for _, v := range c.ShortUrlBlackList {
		m[v] = struct{}{}
	}
	//初始化布隆过滤器
	store := redis.MustNewRedis(c.BloomRedis)
	bitSet := bloom.New(store, "bloom_filter", 20*(1<<20))
	//加载已有的短链接数据(基于内存方式）
	return &ServiceContext{
		Config:         c,
		ShortUrlModel:  model.NewShortUrlMapModel(conn, c.CacheRedis),
		Sequence:       sequence.NewMySQL(c.SequenceDB.DSN),
		ShortBlackList: m, //把 短链接黑名单 放到map
		Filter:         bitSet,
	}
}

// 加载已有的短链接数据 至布隆过滤器中、
/*
func loadData2Bloom(ctx *ServiceContext) error {
	//从数据库中加载数据
	list, err := ctx.ShortUrlModel.FindAll()
	if err != nil {
		return err
	}
	//把数据加载到布隆过滤器中
	for _, v := range list {
		ctx.Filter.Add([]byte(v.Surl.String))
	}
	return nil
}
*/
