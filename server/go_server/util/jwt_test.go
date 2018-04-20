package util

import (
	"testing"
	"time"
)

func TestJSONWebToken(t *testing.T) {
	token := NewJSONWebToken("abc")
	token.Iss = "jd"
	token.Iat = time.Now().Unix()
	token.Exp = time.Now().Add(time.Hour).Unix()
	token.FromUser = 1
	s, err := token.Token()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)

	ok, err := token.Verify(s)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("bad token")
	}
	t.Log(token)
}
