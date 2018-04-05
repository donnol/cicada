package util

import (
	"log"
	"os"
)

var _ = func() error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)
	return nil
}()
