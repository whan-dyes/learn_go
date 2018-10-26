``` go
set GOPATH=e:/go_test
 echo $GOPATH
 go get github.com/astaxie/beego

设置环境变量：GOROOT,GOPATH
在PATH环境变量后追加 $GOROOT/bin，$GOPATH/bin（如：d:\go\bin;e:\go\bin)
可能需要重启使设置生效。
go version 可查看当前Go语言版本
go env 可查看与Go语言相关的环境变量
工作目录下要求创建三个目录，分别为 src，pkg，bin，其中所有源码必须放在src目录下。

rm删除一个目录或文件后，提交时用：
git commit -am  "kkkkkkkkk"
```
