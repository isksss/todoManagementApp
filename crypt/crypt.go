package crypt

import (
	"crypto/sha512"
	"encoding/hex"
)

func passwdEncrypt(passwd string) string {

	hashed := "" // ハッシュの初期化

	sha512 := sha512.Sum512([]byte(passwd))
	hashed = hex.EncodeToString(sha512[:])

	return hashed
}
