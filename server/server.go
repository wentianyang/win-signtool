package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"syscall"

	"github.com/yangtianwen/win-signtool/sign"
)

// 使用 gin 框架,创建一个文件上传服务
// 上传文件到指定目录
// 上传成功后,执行签名

func Server() {

	http.Handle("/", http.FileServer(http.Dir("./client")))

	http.HandleFunc("/upload", uploadFile)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error Starting Server:", err)
		return
	}
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// 解析表单
	// r.ParseMultipartForm(10 << 20) // 10MB

	// 获取文件句柄
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 输出文件信息
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Type: %+v\n", handler.Header.Get("Content-Type"))

	// 创建文件
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error Saving File")
		fmt.Println(err)
		return
	}

	// 获取文件句柄对应的文件描述符
	fd := f.Fd()

	// 获取文件锁
	lock := &syscall.Flock_t{Type: syscall.F_WRLCK}
	if err := syscall.FcntlFlock(fd, syscall.F_SETLK, lock); err != nil {
		fmt.Println("Error Acquiring File Lock")
		fmt.Println(err)
		f.Close()
		return
	}

	// 写入文件到服务器
	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Println("Error Writing File Content")
		fmt.Println(err)
		f.Close()
		return
	}

	// 关闭文件
	f.Close()

	fmt.Fprintf(w, "Successfully Uploaded File\n")

	// 执行签名
	sign.Sign(handler.Filename, func(filename string) {
		fmt.Fprintf(w, "Successfully Signed File: %s\n", filename)
	})
}