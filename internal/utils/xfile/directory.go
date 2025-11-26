package xfile

import (
	"errors"
	"os"
)

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Mkdir(path string) error {
	b, err := PathExists(path)
	if err != nil {
		return nil
	}
	if !b {
		return os.Mkdir(path, os.ModePerm)
	} else {
		return nil
	}
}
