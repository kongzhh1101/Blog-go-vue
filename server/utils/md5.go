package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(src []byte) string {
	m := md5.New()
	m.Write(src)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
