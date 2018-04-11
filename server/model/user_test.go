package model

import (
	"testing"
)

func TestPhoneRegisterCode(t *testing.T) {
	for _, tc := range []string{
		"12345678901",
	} {
		code, err := PhoneRegisterCode(tc)
		if err != nil {
			t.Fatal(err)
		}
		if code == "" {
			t.Fatalf("test failed, got: %v", code)
		}
	}
}
