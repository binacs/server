package base

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"time"

	//"github.com/satori/go.uuid"

	"github.com/BinacsLee/server/types"
)

func GenToken(id string, t time.Duration) string {
	expireTime := time.Now().Add(t).Unix()
	buffer := NewBuffer().Append(expireTime).Append(id).Append(types.TokenSalt)
	md5 := Md5(buffer.String())
	buffer = NewBuffer().Append(expireTime).Append(md5).Append("id").Append(id)
	return buffer.String()
}

func Md5(v string) string {
	ret := md5.New()
	io.WriteString(ret, v)
	return fmt.Sprintf("%x", ret.Sum(nil))
}

func UniqueID() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}

func UniquePWD() string {
	b := make([]byte, 96)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}
