package intlib

import (
	"fmt"
	"strings"
)

// SplitByInt64Slice int64切片转换为字符串
func SplitByInt64Slice(params []int64, separator string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(params), "[]"), " ", separator, -1)
}

// Int64SliceUniq int64切片去重
func Int64SliceUniq(addrs []int64) []int64 {
	result := make([]int64, 0, len(addrs))
	temp := map[int64]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
