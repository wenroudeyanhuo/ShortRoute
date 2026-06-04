// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"ShotTener/internal/svc"
	"ShotTener/internal/types"
	"ShotTener/model"
	"ShotTener/pkg/base62"
	"ShotTener/pkg/connect"
	"ShotTener/pkg/md5"
	"ShotTener/pkg/urltool"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// convert 将长链接转换为短链接
func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// todo: add your logic here and delete this line
	//1校验输入数据->  数据不能为空 链接必须能请求通的网址  判断是否已经转链过  输入不能是一个短链接
	//1.1数据不能为空
	//if len(req.LongURL) == 0 {return nil, errorx.NewDefaultError("请输入要转换的长链接") }
	//使用validator 包来做参数校验
	//1.2链接必须能请求通的网址
	//http.Get(req.LongURL)
	if !connect.Get(req.LongURL) {
		return nil, errors.New("无效的链接")
	}
	//1.3  判断是否已经转链过
	//不是直接查，数据量大，我们这里走md5的索引
	//根据长链接生成md5
	md5Value := md5.Sum([]byte(req.LongURL))
	//拿md5去查数据库是否存在
	//但是svcctx还没有关于数据查询的服务   ->去修改type ServiceContext struct 然后添加连接，到时候主函数启动时会调用
	//使用 model 来查询
	u, errs := l.svcCtx.ShortUrlModel.FindOneByMd5(l.ctx, sql.NullString{String: md5Value, Valid: true})
	//如果不是不存在，说明已经转链过了
	if errs != sqlx.ErrNotFound {
		if errs == nil {
			return nil, fmt.Errorf("该链接已经转换过了%s", u.Surl.String)
		}
		//是真错了
		logx.Errorw("ShortUrlModel.FindOneByMd5 failed", logx.LogField{Key: "error", Value: errs.Error()})
		return nil, errs
	}
	//1.4 输入不能是一个短链接
	basePath, err := urltool.GetBasePath(req.LongURL)
	if err != nil {
		logx.Errorw("url.Parse failed", logx.LogField{
			Key: "lurl", Value: req.LongURL,
		}, logx.LogField{Key: "error", Value: err.Error()})
		return nil, err
	}
	_, err = l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{String: basePath, Valid: true})
	if err != sqlx.ErrNotFound {
		if err == nil {
			return nil, errors.New("该链接已经是短链了")
		}
		//查询错误
		logx.Errorw("ShortUrlModel.FindOneBySurl failed", logx.LogField{
			Key: "error", Value: err.Error(),
		})
		return nil, err
	}
	var short string
	for {
		//2取号  基于mysql实现取号器
		//每来一个转链请求，我们呢使用REPLACE INTO 语句往sequence 表插入一条数据  并且去除主键id作为号码
		seq, err := l.svcCtx.Sequence.Next() //seq就是号
		if err != nil {
			logx.Errorw("Sequence.Next failed", logx.LogField{
				Key: "err", Value: err.Error(),
			})
			return nil, err
		}
		//3号码转短链
		//3.1 安全性
		//3.2 黑名单  避免某些特殊词  例如你可能使用的路由**/health /login等
		short = base62.Int2String(seq)
		//fmt.Println(short)
		if _, ok := l.svcCtx.ShortBlackList[short]; !ok {
			//不在黑名单中跳出循环
			break
		}
	}
	//4存储长短链接映射关系
	if _, err := l.svcCtx.ShortUrlModel.Insert(
		l.ctx,
		&model.ShortUrlMap{
			Md5:  sql.NullString{String: md5Value, Valid: true},
			Surl: sql.NullString{String: short, Valid: true},
			Lurl: sql.NullString{String: req.LongURL, Valid: true},
		},
	); err != nil {
		//	插入有问题
		logx.Errorw("ShortUrlModel.Insert failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}
	//向布隆过滤器中加入新生成的短链接
	err = l.svcCtx.Filter.Add([]byte(short))
	if err != nil {
		logx.Errorw("BloomFilter.Add failed", logx.LogField{Key: "err", Value: err.Error()})
	}
	//5返回响应
	//返回的是短域名+短链接
	shortURL := l.svcCtx.Config.ShortDomain + "/" + short
	return &types.ConvertResponse{ShortURL: shortURL}, nil
}
