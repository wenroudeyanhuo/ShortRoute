// @program:     ShotTener
// @file:        md5.go
// @author:      16574
// @create:      2025-12-14 15:23
// @description:

package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// Sum 对传入的数据进行md5计算
func Sum(data []byte) string {
	h := md5.New()
	h.Write(data)
	//	返回md5 子符串  MD5 固定输出 128 位（16 字节）
	return hex.EncodeToString(h.Sum(nil)) //32位 16进制数
}
