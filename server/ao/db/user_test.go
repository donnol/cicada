package db

import (
	"testing"
)

func TestExistPhone(t *testing.T) {
	for _, tc := range []string{
		"12345678901",
	} {
		exist, err := ExistPhone(tc)
		if err != nil {
			t.Fatal(err)
		}
		if exist != false {
			t.Fatalf("test failed, got: %v", exist)
		}
	}
}
