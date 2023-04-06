package main

import (
	"fmt"
	"net/http"
)

// 使用 gin 框架,创建一个文件上传服务
// 上传文件到指定目录
// 上传成功后,执行签名

func server() {
	http.HandleFunc("/upload", uploadFile)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error Starting Server:", err)
		return
	}
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// 解析表单
	r.ParseMultipartForm(10 << 20) // 10MB

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

	// 写入文件到服务器
	// TODO: 替换成您自己的文件保存逻辑
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}