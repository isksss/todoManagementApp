package crypt

import (
	"crypto/sha512"
	"encoding/hex"
	"math/rand"
	"time"
)

func PasswdEncrypt(passwd string) (hashed string) {

	hashed = "" // ハッシュの初期化

	sha512 := sha512.Sum512([]byte(passwd))
	hashed = hex.EncodeToString(sha512[:])

	return
}

func CreateRandomInt(min int, max int) (num int) {

	rand.Seed(time.Now().UnixNano())
	num = min + rand.Intn(max-min)

	return
}

func NewSessionKey() {

}
