// @program:     ShotTener
// @file:        sequence.go
// @author:      16574
// @create:      2025-12-15 11:45
// @description:

package sequence

// Sequence 取号器接口
type Sequence interface {
	Next() (uint64, error)
}
