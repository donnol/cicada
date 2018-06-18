package model

import (
	"testing"
)

func TestNewFile(t *testing.T) {
	kind := "file"
	min, max := 0, 1024*1024*64
	dir := "/home/jd/tmp"
	format := []string{"jpg", "png", "gif"}
	file := NewFile(SetKind(kind), SetRange(min, max), SetFormat(format), SetDir(dir))
	if file.kind != kind ||
		file.min != min ||
		file.max != max ||
		file.dir != dir ||
		len(file.formats) != len(format) ||
		file.formats[0] != format[0] ||
		file.formats[1] != format[1] ||
		file.formats[2] != format[2] {
		t.Fatalf("bad file, %+v", file)
	}
}
