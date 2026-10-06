// flex_int64.go 协议响应容错类型。
// 协议服务的整数字段在不同场景可能返回 number 或 string，统一用 FlexInt64 兼容两种格式。

package sdk_ipad

import (
	"math"
	"strconv"
	"strings"
)

// FlexInt64 同时兼容协议返回的 JSON number 与字符串数字两种格式的整数。
// 使用示例：
//
//	type MainData struct {
//		BanFinishTime FlexInt64 `json:"ban_finish_time"`
//	}
type FlexInt64 int64

// UnmarshalJSON 兼容解析 number（1746252388）、字符串数字（"1746252388"）、
// null 与空字符串（均解析为 0）。
func (f *FlexInt64) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	if s == "" || s == "null" {
		*f = 0
		return nil
	}
	n, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*f = FlexInt64(n)
		return nil
	}
	// 兼容 JSON number 整数浮点格式（如 1746252388.0），其余非法值返回 error
	fv, ferr := strconv.ParseFloat(s, 64)
	if ferr != nil || fv != math.Trunc(fv) || fv >= 9.223372036854776e18 || fv <= -9.223372036854776e18 {
		return err
	}
	*f = FlexInt64(fv)
	return nil
}

// MarshalJSON 统一按 number 格式序列化，与协议保持一致。
func (f FlexInt64) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(f), 10)), nil
}
