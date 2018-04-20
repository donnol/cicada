package util

import (
	"math/rand"
	"time"
)

var _ = func() error {
	rand.Seed(time.Now().Unix())
	return nil
}
