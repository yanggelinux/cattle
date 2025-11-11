package util

import (
	"context"
	"crypto/rand"
	"github.com/google/uuid"
	"math/big"
	"regexp"
	"strings"
	"unicode"
)

// SnakeString
/**
 * 驼峰转蛇形 snake string
 * @description XxYy to xx_yy , XxYY to xx_y_y
 * @date 2020/7/30
 * @param s 需要转换的字符串
 * @return string
 **/
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

func RemoveDupString(arr []string) []string {
	//字符串切片去重复
	if len(arr) == 0 {
		return arr
	}
	result := make([]string, 0, len(arr))
	temp := map[string]struct{}{}
	for _, item := range arr {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// DeleteExtraSpace 多个空格转为一个空格
func DeleteExtraSpace(s string) string {
	regstr := "\\s{2,}"
	reg, _ := regexp.Compile(regstr)
	tmpstr := make([]byte, len(s))
	copy(tmpstr, s)
	spc_index := reg.FindStringIndex(string(tmpstr))
	for len(spc_index) > 0 {
		tmpstr = append(tmpstr[:spc_index[0]+1], tmpstr[spc_index[1]:]...)
		spc_index = reg.FindStringIndex(string(tmpstr))
	}
	return string(tmpstr)

}

func GetUserName(ctx context.Context) string {
	return ToString(ctx.Value("X-Username"))
}
func GetUserID(ctx context.Context) string {
	return ToString(ctx.Value("X-UserID"))
}
func GetRequestID(ctx context.Context) string {
	return ToString(ctx.Value("X-RequestID"))
}
func GetAuthorization(ctx context.Context) string {
	return ToString(ctx.Value("X-Authorization"))
}
func GetSuper(ctx context.Context) string {
	return ToString(ctx.Value("X-Super"))
}
func GetDeptName(ctx context.Context) string {
	return ToString(ctx.Value("X-DeptName"))
}

// StrHasHan 判断字符串中是否有汉字
func StrHasHan(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

func RandString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		// 生成 [0, len(letters)) 范围内的随机整数
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return ""
		}
		result[i] = letters[idx.Int64()]
	}
	return string(result)
}

func genUUID() string {
	return uuid.New().String()
}
func GenUUIDv4() string {
	u, err := uuid.NewRandomFromReader(rand.Reader)
	if err != nil {
		return genUUID()
	}
	return u.String()
}
