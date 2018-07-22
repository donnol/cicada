package util

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var _ = func() error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 每天新建一个日志文件
	timeLayout := "2006-01-02"
	fileName := time.Now().Format(timeLayout)
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	path, err := filepath.Abs(file)
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(path)
	fullPath := filepath.Join(dir, "log", fileName+".log")

	f, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Println(err)
		return err
	}
	log.SetOutput(f)

	return nil
}()
