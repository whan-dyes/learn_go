package function

import (
	. "fmt"
)

func MyDefer() {
	Println("<<<<<<<<<<<<<<<<<<<<< MyDefer >>>>>>>>>>>>>>>>>>>>>>>>")
	defer1()
	defer2()
	defer3()
	defer4()
}

func defer1() {
	Println("............ defer1() ...........")
}

func defer2() {
	Println("............ defer2() ...........")
	defer func() {
		if err := recover(); err != nil {
			Println("Recover in defer2")
		}
	}()
	panic("Panic in defer2")
}

func defer3() {
	Println("............ defer3() ...........")
}

func defer4() {
	Println("............ defer4() ...........")
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		defer Println("defer i=", i)
		defer func() { Println("defer_closure i=", i) }()
		fs[i] = func() { Println("closure i=", i) }
	}

	for _, f := range fs {
		f()
	}
}

/*
defer

的执行方式类似其它语言中的析构函数，在函数体执行结束后
按照调用顺序的相反顺序逐个执行
即使函数发生严重错误也会执行
支持匿名函数的调用
常用于资源清理、文件关闭、解锁以及记录时间等操作
通过与匿名函数配合可在return之后修改函数计算结果
如果函数体内某个变量作为defer时匿名函数的参数，则在定义defer
时即已经获得了拷贝，否则则是引用某个变量的地址

Go 没有异常机制，但有 panic/recover 模式来处理错误
Panic 可以在任何地方引发，但recover只有在defer调用的函数中有效
*/
