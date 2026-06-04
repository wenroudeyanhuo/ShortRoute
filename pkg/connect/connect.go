// @program:     ShotTener
// @file:        connect.go
// @author:      16574
// @create:      2025-12-14 14:47
// @description:

package connect

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"time"
)

// client 全局的http客户端
var client = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
	Timeout: 2 * time.Second,
}

// Get 判断url是否能请求通
func Get(url string) bool {
	resp, err := client.Get(url)
	if err != nil {
		logx.Errorw("connect client.Get failed", logx.LogField{Key: "err", Value: err.Error()})
		return false
	}
	//读取完效应体后应该关闭
	resp.Body.Close()
	//并不会清除数据，只是将数据从缓冲区中读取到程序中
	//别人发一个跳转链接也不会给过
	return resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusBadRequest

}
