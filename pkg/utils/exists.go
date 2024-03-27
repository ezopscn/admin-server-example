package utils

import (
	"fmt"
	"os"
)

// 判断文件是否存在
func FileExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if !os.IsNotExist(err) {
		if !stat.IsDir() {
			return true, nil
		}
		return false, fmt.Errorf("路径 %s 已经存在, 但它是一个目录", path)
	}
	return false, fmt.Errorf("路径 %s 不存在", path)
}

// 判断目录是否存在
func DirExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if !os.IsNotExist(err) {
		if stat.IsDir() {
			return true, nil
		}
		return false, fmt.Errorf("路径 %s 已经存在, 但它是一个文件", path)
	}
	return false, fmt.Errorf("路径 %s 不存在", path)
}
