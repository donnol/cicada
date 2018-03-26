package ao

import (
	"cicada/server/ao/db"
	"testing"
)

func TestRegisterCode(t *testing.T) {
	for _, tc := range []string{
		"12345678901",
	} {
		code, err := RegisterCode(tc)
		if err != nil {
			t.Fatal(err)
		}
		exist, err := db.ExistPhoneCode(tc, code)
		if err != nil {
			t.Fatal(err)
		}
		if exist != true {
			t.Fatalf("test failed, got: %v", exist)
		}
	}
}
