package sign

import (
	"os"
	"os/exec"
	"path/filepath"
)

func Sign(dirPath string, result func(error, string)) {
	err := walkDir(dirPath, func(filename string) {
		// 打印文件名
		println(filename)
		// 执行签名
		sha1SignFile(filename, result)
		sha256SignFile(filename, result)
	})
	if err != nil {
		result(err, "")
	}
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
func sha1SignFile(filename string, result func(error, string)) {
	sha1Cmd := exec.Command("signtool", "sign", "/n", "公司名称", "/fd", "SHA1", "/t", "http://timestamp.digicert.com", filename)
	msg, err := sha1Cmd.CombinedOutput()
	if err != nil {
		result(err, "SHA1签名失败: "+err.Error()+" "+string(msg))
		return
	}
	result(nil, "SHA1签名成功: "+filename)
}

// SHA256签名
func sha256SignFile(filename string, result func(error, string)) {
	sha256Cmd := exec.Command("signtool", "sign", "/n", "公司名称", "/as", "/fd", "SHA256", "/tr", "http://timestamp.digicert.com", "/td", "SHA256", filename)
	msg, err := sha256Cmd.CombinedOutput()
	if err != nil {
		result(err, "SHA256签名失败: "+err.Error()+" "+string(msg))
		return
	}
	result(nil, "SHA256签名成功: "+filename)
}
