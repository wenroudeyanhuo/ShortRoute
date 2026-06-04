// @program:     ShotTener
// @file:        urltool.go
// @author:      16574
// @create:      2025-12-15 10:09
// @description:

package urltool

import (
	"net/url"
	"path"
)

// GetBasePath 获取url的最后一节
func GetBasePath(Url string) (string, error) {
	//输入的是一个完整的url qimi.cn/123456
	myUrl, err := url.Parse(Url)
	if err != nil {
		return "", err
	}
	basePath := path.Base(myUrl.Path)
	return basePath, nil
}
