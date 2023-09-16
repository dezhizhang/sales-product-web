package utils

import (
	"go.uber.org/zap"
	"strconv"
)

// StringToNumber 字符串类型转换成int
func StringToNumber(value string) int {
	number, err := strconv.Atoi(value)
	if err != nil {
		zap.S().Errorw("类型转换错误%s", err.Error())
	}
	return number
}
