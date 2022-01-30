package logging

import (
	"os"
	"path/filepath"
)

func CreateFile(pathName string) (*os.File, error) {
	dir := filepath.Dir(pathName)
	// 文件 目录 权限
	err := os.MkdirAll(dir, 0775)
	if err != nil {
		return nil, err
	}
	f, err := os.OpenFile(pathName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return f, nil
}
