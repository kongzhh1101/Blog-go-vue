package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 转换为哈希值

func MD5(src []byte) string {
	m := md5.New()
	m.Write(src)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
