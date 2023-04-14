package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"testing"
)

func Test_hmac(t *testing.T) {

	h := hmac.New(sha256.New, []byte("CsYyCc=="))

	h.Write([]byte("GET"))
	h.Write([]byte("/v1/verify"))
	h.Write([]byte("rQJEk4mzg"))
	h.Write([]byte("123"))
	h.Write([]byte("123"))
	// 写入 nil 和 不调用的结果一样
	h.Write(nil)

	fmt.Println(base64.StdEncoding.EncodeToString(h.Sum(nil)))

}
