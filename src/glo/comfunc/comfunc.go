package comfunc

import (
	"fmt"
	"mos/src/glo/encrypt"
	"strings"
	"time"
)

// DefaultPassword const Params
const DefaultPassword string = "123456"

// TokenInvaild Const Params
const TokenInvaild string = "Token invaild"

// FormatTs 格式化时间戳为字符串
func FormatTs(ts int64) (res string) {
	res = time.Unix(ts, 0).Format("2006-01-02 15:04:05")
	return
}

// FlorPageInt 分页整除判断
func FlorPageInt(pageSize int, total int) (res int) {
	res = 0
	if pageSize <= 0 {
		return
	}
	mix := total % pageSize
	if mix != 0 {
		res = total/pageSize + 1
	} else {
		res = total / pageSize
	}
	return
}

// GetDefaultPage 校验页数，小于0返回默认1
func GetDefaultPage(page int) (res int) {
	if page <= 0 {
		res = 1
		return
	}
	res = page
	return
}

// GetDefaultPageSize 校验每页行数，小于0返回默认20
func GetDefaultPageSize(pageSize int) (res int) {
	if pageSize <= 0 {
		res = 20
		return
	}
	res = pageSize
	return
}

// EncryptToken Init Token
func EncryptToken(srcStr string, timeTs int64, key string) (token string) {
	encStr := fmt.Sprintf("%s_%d", srcStr, timeTs)
	token, _ = encrypt.AesEncryptString([]byte(encStr), []byte(key))
	return
}

// DecryptToken Token
func DecryptToken(token string, key string) (srcStr string) {
	encStr, _ := encrypt.AesDecryptString(token, []byte(key))
	srcStr = strings.Split(string(encStr), "_")[0]
	return
}
