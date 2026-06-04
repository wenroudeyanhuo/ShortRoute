// @program:     ShotTener
// @file:        base62.go
// @author:      16574
// @create:      2025-12-15 12:11
// @description:

package base62

import (
	"math"
	"strings"
)

const defaultBaseStr = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 62进制转换
// 0-9 -> 0-9,a-z -> 10-35    A-Z 36->61
// 实现 62进制转换
// const base62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
//const base62 = "pqrKLMNOP4stu012abchCDEijkl389mnovwxyzABFGdefgHIJ567QRSTUVWXYZ"

//为了避免被人恶意请求，我们可以将上面的字符串打乱
/*
//这种写法实现了递归递归，但是效率不高

	func Int2String(seq uint64) string {
		if seq == 0 {
			return string(base62[0])
		}

		return Int2String(seq/62) + string(base62[seq%62])

}
*/

// 这两个字段将从外面获取
var (
	baseStr    = defaultBaseStr
	baseStrlen = uint64(len(defaultBaseStr))
)

// MustInit 要使用base62 这个包就必须进行初始化
func MustInit(bs string) {
	if len(bs) == 0 {
		panic("base string can not be empty")
	}
	if len(bs) != 62 {
		panic("base string length error")
	}
	baseStr = bs
	baseStrlen = uint64(len(baseStr))
}

// Int2String 十进制数转为62进制子符串
func Int2String(seq uint64) string {
	if seq == 0 {
		return string(baseStr[0])
	}
	bl := []byte{}
	for seq > 0 {
		bl = append(bl, baseStr[seq%62])
		seq /= 62
	}
	return string(reverse(bl))
}

// 62进制字符串转换为十进制
func String2Int(seq string) (result uint64) {
	bl := []byte(seq)
	bl = reverse(bl)
	for idx, b := range bl {
		base := math.Pow(62, float64(idx))
		result += uint64(strings.Index(baseStr, string(b))) * uint64(base)
	}
	return result
}
func reverse(str []byte) []byte {
	for i, j := 0, len(str)-1; i < len(str)/2; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return str
}
