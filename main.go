package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	fmt.Println("window signtool start....")

	file := os.Args[1]

	if file == "" {
		log.Fatalf("请输入需要签名的文件或者文件夹")
	}

	fmt.Printf("签名的文件=====> %s\n", file)

	cmdSHA1 := exec.Command(
		"signtool",
		"sign",
		"/n",
		"\"公司\"",
		"/t",
		"http://timestamp.digicert.com",
		"/fd",
		"SHA1",
		file)

	cmdSHA256 := exec.Command(
		"signtool",
		"sign",
		"/n",
		"\"公司的名字\"",
		"/tr",
		"http://timestamp.digicert.com",
		"/td",
		"SHA256",
		file,
	)

	infoSHA1, errSHA1 := cmdSHA1.CombinedOutput()

	infoSHA256, errSHA256 := cmdSHA256.CombinedOutput()

	if errSHA256 != nil {
		log.Fatalf("sign sha256 error: %s\n", infoSHA256)
	}

	if errSHA1 != nil {
		log.Fatalf("signtool error with %s\n", infoSHA1)
	}

	fmt.Printf("sign SHA1 result %s\n", infoSHA1)
	fmt.Println("--------------------------------------------------------------------")
	fmt.Printf("sign SHA256 result %s\n", infoSHA256)
}
