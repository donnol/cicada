package model

import (
	"errors"
	"io"
)

// FileOption 配置方法
type FileOption func(*File)

// File 文件模块
type File struct {
	// 初始化设置
	kind    string   // 存储类型，db表示数据库，file表示文件，link表示如七牛的第三方存储；默认是file
	max     int      // 最大大小
	min     int      // 最小大小
	dir     string   // 存储路径
	formats []string // 支持的类型，如jpg，gif，png，csv等，如果包含'*'，表示类型不限
}

// NewFile 新建文件
func NewFile(opts ...FileOption) *File {
	file := &File{}
	for _, opt := range opts {
		opt(file)
	}

	return file
}

// SetKind 设置类型
func SetKind(kind string) FileOption {
	if kind != "db" &&
		kind != "file" &&
		kind != "link" {
		panic("wrong kind")
	}
	return func(f *File) {
		f.kind = kind
	}
}

// SetRange 设置大小范围
func SetRange(min, max int) FileOption {
	return func(f *File) {
		f.min = min
		f.max = max
	}
}

// SetDir 设置存储目录
func SetDir(dir string) FileOption {
	return func(f *File) {
		f.dir = dir
	}
}

// SetFormat 设置文件格式
func SetFormat(formats []string) FileOption {
	return func(f *File) {
		f.formats = formats
	}
}

// Upload 上传
func (f *File) Upload(rc io.ReadCloser, size int, format string) (id int, err error) {
	// 如果是异步的话，就需要由被调用方来决定什么时候关闭
	// 如果是同步的话，调用方自己关闭也可以
	defer rc.Close()

	// 检查参数
	if err = f.check(size, format); err != nil {
		return
	}

	// TODO
	switch f.kind {
	case "db":
	case "link":
	case "file":
		fallthrough
	default:
		// 写表，记录文件hash，大小，路径/链接，创建人，创建时间

		// 按文件id对100的余数保存文件到相应的目录
	}

	return
}

// Download 下载
func (f *File) Download(id int) (rc io.ReadCloser, err error) {
	// TODO
	// 通过id获取文件路径/链接

	// 打开文件

	return
}

func (f *File) check(size int, format string) (err error) {
	if size < f.min || size > f.max {
		err = errors.New("size too small or too big")
		return
	}
	var validFormat bool
	for _, form := range f.formats {
		if form == "*" || form == format {
			validFormat = true
			break
		}
	}
	if !validFormat {
		err = errors.New("invalid format")
		return
	}

	return
}
