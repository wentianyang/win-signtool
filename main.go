package main

import (
	"github.com/yangtianwen/win-signtool/server"
)

func main() {
	// // 获取当前目录
	// dir, _ := os.Getwd()
	// // 遍历目录
	// walkDir(dir, func(filename string) {
	// 	// 打印文件名
	// 	println(filename)
	// 	// 执行 sha1 签名
	// 	sha1SignFile(filename)
	// 	// 执行 sha256 签名
	// 	sha256SignFile(filename)
	// })

	// 启动服务
	server.Server()
}
