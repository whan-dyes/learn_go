```go


运行不要退出，将输出到本地8080端口，然后在网页查看 localhost:8080,即可查看官网go文档。
$ godoc -http=:8080
//----------------------------------
 git clone https://github.com/golang/tools.git

 通过下面命令安装gotags
  go get -u github.com/jstemmer/gotags
 在你的源码目录下执行下面命令声称ctags：
  gotags -tag-relative=true -R=true -sort=true -f="tags" -fields=+l .
 这将会在你的项目根目录下生成新的tags文件，让你的代码提示更加智能

//-------------------------------
Go Web开发框架
  Beego（http://beego.me/)
  MVC框架，作者中国人，框架中文文档丰富，便于初学者

  Revel （http://revel.github.io/)
  思路来自Java的PlayerFramework，相对Beego难一点。

  Martini （http://martini.codegangsta.io/)
简单灵活，大量的使用反射，初学不易上手。
```
