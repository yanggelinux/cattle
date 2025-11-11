package util

import (
	"regexp"
)

var VerRuleMapping = map[int8]string{
	1:  "无规则",
	2:  "仅数字",
	3:  "仅英文字母",
	4:  "仅小写英文字母",
	5:  "仅大写英文字母",
	6:  "仅英文字母和数字",
	7:  "仅英文字母、数字和下划线",
	8:  "仅英文字母、数字和中横线",
	9:  "仅中文",
	10: "邮件",
	11: "手机号",
	12: "IP地址",
	13: "时间字符串,YYYY-mm-dd HH:MM:SS",
}

// 仅数字
func IsOnlyDigits(s string) bool {
	return regexp.MustCompile(`^\d+$`).MatchString(s)
}

// 仅英文字母
func IsOnlyLetters(s string) bool {
	return regexp.MustCompile(`^[A-Za-z]+$`).MatchString(s)
}

// 仅小写英文字母
func IsOnlyLowercaseLetters(s string) bool {
	return regexp.MustCompile(`^[a-z]+$`).MatchString(s)
}

// 仅大写英文字母
func IsOnlyUppercaseLetters(s string) bool {
	return regexp.MustCompile(`^[A-Z]+$`).MatchString(s)
}

// 仅英文字母和数字
func IsOnlyLettersAndDigits(s string) bool {
	return regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString(s)
}

// 仅英文字母、数字和下划线
func IsOnlyLettersDigitsUnderscore(s string) bool {
	return regexp.MustCompile(`^\w+$`).MatchString(s)
}

// 仅英文字母、数字和中横线
func IsOnlyLettersDigitsHyphen(s string) bool {
	return regexp.MustCompile(`^[A-Za-z0-9-]+$`).MatchString(s)
}

// 仅中文
func IsOnlyChinese(s string) bool {
	return regexp.MustCompile(`^[\p{Han}]+$`).MatchString(s)
}

// 邮件格式
func IsEmail(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`).MatchString(s)
}

// 中国大陆手机号格式
func IsChineseMobile(s string) bool {
	return regexp.MustCompile(`^1[3-9]\d{9}$`).MatchString(s)
}

// IPv4 地址格式
func IsIPv4(s string) bool {
	return regexp.MustCompile(`^(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)){3}$`).MatchString(s)
}

// 时间字符串格式：YYYY-MM-DD HH:MM:SS
//func IsDatetimeString(s string) bool {
//	return regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`).MatchString(s)
//}
//
//var datetimeRegex = regexp.MustCompile(
//	`^(19|20)\d{2}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01]) ` + // 年月日
//		`(0\d|1\d|2[0-3]):([0-5]\d):([0-5]\d)$`, // 时分秒
//)

// 校验日期时间格式：YYYY-MM-DD HH:MM:SS
func IsDatetimeString(s string) bool {
	pattern := `^(19|20)\d{2}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01]) ` + // 年月日
		`(0\d|1\d|2[0-3]):([0-5]\d):([0-5]\d)$`
	return regexp.MustCompile(pattern).MatchString(s)
}

// 校验日期字符串格式：YYYY-MM-DD
func IsDateString(s string) bool {
	// 年：1000–9999
	// 月：01–12
	// 日：01–31（这里用通用校验，不验证闰年等逻辑）
	pattern := `^(19|20)\d{2}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`
	return regexp.MustCompile(pattern).MatchString(s)
}

// 交完表单的值
func CheckValue(rule int8, val interface{}) bool {

	switch rule {
	case 1:
		return true
	case 2:
		switch v := val.(type) {
		case int, int8, int16, int32, int64:
			return true
		case float32, float64:
			return true
		case string:
			if !IsOnlyDigits(v) {
				return false
			}
			return true
		default:
			return false
		}
	case 3:
		v, ok := val.(string)
		if ok {
			if !IsOnlyLetters(v) {
				return false
			}
			return true
		} else {
			return false
		}
	case 4:
		v, ok := val.(string)
		if ok {
			if !IsOnlyLowercaseLetters(v) {
				return false
			}
			return true
		} else {
			return false
		}
	case 5:
		v, ok := val.(string)
		if ok {
			if !IsOnlyUppercaseLetters(v) {
				return false
			}
			return true
		} else {
			return false
		}
	case 6:
		v, ok := val.(string)
		if ok {
			if !IsOnlyLettersAndDigits(v) {
				return false
			}
			return true
		} else {
			return false
		}
	case 7:
		v, ok := val.(string)
		if ok {
			if !IsOnlyLettersDigitsUnderscore(v) {
				return false
			}
			return true
		} else {
			return false
		}
	case 8:
		v, ok := val.(string)
		if ok {
			if !IsOnlyLettersDigitsHyphen(v) {
				return false
			}
			return true
		} else {
			return false
		}
	case 9:
		v, ok := val.(string)
		if ok {
			if !IsOnlyChinese(v) {
				return false
			}
			return true
		} else {
			return false
		}
	case 10:
		v, ok := val.(string)
		if ok {
			if !IsEmail(v) {
				return false
			}
			return true
		} else {
			return false
		}
	case 11:
		v, ok := val.(string)
		if ok {
			if !IsChineseMobile(v) {
				return false
			}
			return true
		} else {
			return false
		}
	case 12:
		v, ok := val.(string)
		if ok {
			if !IsIPv4(v) {
				return false
			}
			return true
		} else {
			return false
		}
	case 13:
		v, ok := val.(string)
		if ok {
			if !IsDatetimeString(v) {
				return false
			}
			return true
		} else {
			return false
		}
	case 14:
		v, ok := val.(string)
		if ok {
			if !IsDateString(v) {
				return false
			}
			return true
		} else {
			return false
		}
	default:
		return true
	}
}
