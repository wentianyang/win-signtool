package sign

import (
	"os"
	"os/exec"
	"path/filepath"
)

func Sign(dirPath string, sign func(string)) {
	walkDir(dirPath, func(filename string) {
		// 打印文件名
		println(filename)
		// 执行签名
		sha1SignFile(filename)
		sha256SignFile(filename)
		// 返回信息
		sign(filename)
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

// SHA1签名
func sha1SignFile(filename string) {
	sha1Cmd := exec.Command("signtool", "sign", "/n", "公司名称", "/fd", "SHA1", "/t", "http://timestamp.digicert.com", filename)
	sha1Cmd.CombinedOutput()
}

// SHA256签名
func sha256SignFile(filename string) {
	sha256Cmd := exec.Command("signtool", "sign", "/n", "公司名称", "/as", "/fd", "SHA256", "/tr", "http://timestamp.digicert.com", "/td", "SHA256", filename)
	sha256Cmd.CombinedOutput()
}
