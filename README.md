# PC 签名程序
[SignTool使用说明](https://docs.microsoft.com/zh-cn/dotnet/framework/tools/signtool-exe)
## SHA1 签名

签名命令：

```
signtool sign /n "公司" /t http://timestamp.digicert.com fd SHA1 xxx文件
```

## SHA256 签名
```
signtool sign /n "公司" /tr http://timestamp.digicert.com /td SHA256 xxx文件
```