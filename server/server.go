package server

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/yangtianwen/win-signtool/sign"
)

func Server() {

	http.Handle("/", http.FileServer(http.Dir("./client")))

	http.HandleFunc("/upload", uploadFile)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error Starting Server:", err)
		return
	}
	println("服务开启: http://localhost:8080")
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

	// 写入文件到服务器
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error Saving File")
		fmt.Println(err)
		return
	}
	f.Close()
	io.Copy(f, file)

	fmt.Fprintf(w, "Successfully Uploaded File\n")

	sign.Sign(handler.Filename, func(msg string) {
		fmt.Println(msg)
	})
}
