package base

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"

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
