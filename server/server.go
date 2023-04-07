package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

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
	io.Copy(f, file)
	f.Close()

	fmt.Fprintf(w, "Successfully Uploaded File\n")

	sign.Sign(handler.Filename, func(err error, msg string) {
		fmt.Fprintln(w, msg)

		// 打开签名后的文件
		signedFile, err := os.Open(handler.Filename)
		if err != nil {
			fmt.Println("Error Opening Signed File")
			fmt.Println(err)
			return
		}
		defer signedFile.Close()

		// 将签名后的文件发送给客户端
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(handler.Filename)))
		io.Copy(w, signedFile)
	})
}
