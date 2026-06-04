// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"ShotTener/internal/svc"
	"ShotTener/internal/types"
	"context"
	"database/sql"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	Err404 = errors.New("404")
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	// todo: add your logic here and delete this line
	// 查看短链接 ，输入一个短链接  重定向回真实的连接
	//1.根据短链接查询原始的长连接
	//布隆过滤器
	//不存在的短链接直接返回404，不需要后续处理
	//a.基于内存版本bloom   重启之后就没了，所以每次重启都会重新初始化
	//b.基于redis版本的bloom 一直有 (go-zero自带的)
	//注意，短链接生成时也要在布隆过滤器中存放
	exists, err := l.svcCtx.Filter.Exists([]byte(req.ShortURL))
	if err != nil {
		logx.Errorw("Filter.Exists Failed", logx.Field("error", err.Error()))
	}
	//不存在直接返回
	if !exists {
		return nil, Err404
	}
	//查询数据之前可增加缓存层
	u, err := l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{String: req.ShortURL, Valid: true})
	if err != nil {
		//没有查到
		if err == sql.ErrNoRows {
			return nil, errors.New("短链接不存在")
		}
		//查询出错
		logx.Errorw(" ShortUrlModel.FindOneBySurl Failed 查询短链接失败", logx.Field("err", l.ctx.Value("X-Request-Id")), logx.Field("error", err.Error()))
		return nil, err
	}
	//2.返回重定向响应
	//不在这层做处理，直接返回 交给handler处理
	return &types.ShowResponse{
		LongURL: u.Lurl.String,
	}, nil

}
