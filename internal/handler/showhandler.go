// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"ShotTener/internal/logic"
	"ShotTener/internal/svc"
	"ShotTener/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShowRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		// 使用validator 包来做参数校验 参数规则校验
		if err := validator.New().StructCtx(r.Context(), &req); err != nil {
			logx.Errorw("validator check failed", logx.LogField{Key: "err", Value: err.Error()})
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewShowLogic(r.Context(), svcCtx)
		resp, err := l.Show(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			//这里要改造，返回重定向的响应
			//httpx.OkJsonCtx(r.Context(), w, resp)
			http.Redirect(w, r, resp.LongURL, http.StatusFound)
		}
	}
}
