package util

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strings"
)

// JSONWebToken json 网络令牌
type JSONWebToken struct {
	JSONWebTokenHeader
	JSONWebTokenPayload
	secret string
}

// JSONWebTokenHeader json 网络令牌头部
type JSONWebTokenHeader struct {
	Typ string // 表明这是一个 jwt
	Alg string // 加密方法，如"HS256"
}

// JSONWebTokenPayload json 网络令牌载荷
type JSONWebTokenPayload struct {
	Iss        string // 该JWT的签发者
	Iat        int64  // 签发时间, 时间戳
	Exp        int64  // 过期时间，时间戳
	Aud        string // 接收方
	Sub        string // 面向用户
	FromUser   int    // 发出用户
	TargetUser int    // 目标用户
}

// NewJSONWebToken 创建
func NewJSONWebToken(secret string) *JSONWebToken {
	return &JSONWebToken{
		JSONWebTokenHeader: JSONWebTokenHeader{Typ: "jwt", Alg: "HS256"},
		secret:             secret,
	}
}

// Token 令牌
func (j *JSONWebToken) Token() (s string, err error) {
	str := ""
	for i, v := range []interface{}{
		j.JSONWebTokenHeader,
		j.JSONWebTokenPayload,
	} {
		// 生成 json
		var jsonBytes []byte
		jsonBytes, err = json.Marshal(v)
		if err != nil {
			return
		}
		// base64 encode
		w := new(bytes.Buffer)
		wc := base64.NewEncoder(base64.StdEncoding, w)
		defer wc.Close()
		_, err = wc.Write(jsonBytes)
		if err != nil {
			return
		}
		// 拼接
		str += w.String()
		if i == 0 {
			str += "."
		}
	}
	// hash
	h := hmac.New(sha256.New, []byte(j.secret))
	_, err = h.Write([]byte(str))
	if err != nil {
		return
	}
	s = string(h.Sum(nil))

	s = str + "." + s
	return
}

// Verify 校验
func (j *JSONWebToken) Verify(token string) (ok bool, err error) {
	tokens := strings.Split(token, ".")
	if len(tokens) != 3 {
		return
	}

	// 校验签名
	str := tokens[0] + "." + tokens[1]
	h := hmac.New(sha256.New, []byte(j.secret))
	_, err = h.Write([]byte(str))
	if err != nil {
		return
	}
	s := string(h.Sum(nil))
	ok = s == tokens[2]

	// 解析
	for i := 0; i < 2; i++ {
		r := bytes.NewBuffer([]byte(tokens[0]))
		rc := base64.NewDecoder(base64.StdEncoding, r)
		decodedLen := base64.StdEncoding.DecodedLen(len(tokens[0]))
		data := make([]byte, decodedLen)
		_, err = rc.Read(data)
		if err != nil {
			return
		}
		if i == 0 {
			if err = json.Unmarshal(data, &j.JSONWebTokenHeader); err != nil {
				return
			}
		} else if i == 1 {
			if err = json.Unmarshal(data, &j.JSONWebTokenPayload); err != nil {
				return
			}
		}
	}

	return
}
