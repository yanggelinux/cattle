package util

import (
	"fmt"
	"strconv"
	"strings"
)

func toFloat64(t interface{}) (float64, error) {
	switch v := t.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		return 0, fmt.Errorf("toFloat64: unsupported type %T", t)
	}
}

// ToFloat64 将任意支持的数值类型或字符串转换为 float64
func ToFloat64(t interface{}) float64 {
	f, err := toFloat64(t)
	if err != nil {
		return 0
	}
	return f
}

// ToInteger64 将任意支持的数值类型或字符串转换为 int64
func ToInteger64(t interface{}) int64 {
	f, err := toFloat64(t)
	if err != nil {
		return 0
	}
	return int64(f)
}

// ToUInteger64 将任意支持的数值类型或字符串转换为 uint64
func ToUInteger64(t interface{}) uint64 {
	f, err := toFloat64(t)
	if err != nil {
		return 0
	}
	return uint64(f)
}

// ToInteger 将任意支持的数值类型或字符串转换为 int
func ToInteger(t interface{}) int {
	f, err := toFloat64(t)
	if err != nil {
		return 0
	}
	return int(f)
}

// ToString 将任意基础类型或实现了 fmt.Stringer 的对象格式化为字符串
func ToString(t interface{}) string {
	if t == nil {
		return ""
	}
	switch v := t.(type) {
	case string:
		return v
	case fmt.Stringer:
		return v.String()
	default:
		// 对于基础数值、布尔值、数组、切片、map 等，Sprint 都能给出合理的文本表示
		s := fmt.Sprint(v)
		if s == "<nil>" {
			return ""
		}
		return s
	}
}

func StringToIntSlice(s string) ([]int, error) {
	var sList []int
	ss := strings.Split(strings.TrimLeft(strings.TrimRight(s, "]"), "["), ",")
	for _, s := range ss {
		i, err := strconv.Atoi(s)
		if err != nil {
			return sList, err
		}
		sList = append(sList, i)
	}
	return sList, nil
}

func StringToStringSlice(s string) []string {
	var sList []string
	ss := strings.Split(strings.TrimLeft(strings.TrimRight(s, "]"), "["), ",")
	for _, s := range ss {
		sList = append(sList, s)
	}
	return sList
}
