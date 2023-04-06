package main

import (
	"os"
	"path/filepath"
)

func main() {
	// 获取当前目录
	dir, _ := os.Getwd()
	// 遍历目录
	walkDir(dir, func(filename string) {
		// 打印文件名
		println(filename)
	})
}

// 循环遍历文件夹下的所有文件
func walkDir(dirPth string, f func(string)) error {
	return filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		f(filename)
		return nil
	})
}
