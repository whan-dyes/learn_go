package channel

import (
	. "fmt"
	"runtime"
	"sync"
	"time"
)

/*
并发concurrency

很多人都是冲着 Go 大肆宣扬的高并发而忍不住跃跃欲试，但其实从
源码的解析来看，goroutine 只是由官方实现的超级“线程池”而已。
不过话说回来，每个实例 4-5KB 的栈内存占用和由于实现机制而大幅
减少的创建和销毁开销，是制造 Go 号称的高并发的根本原因。另外，
goroutine 的简单易用，也在语言层面上给予了开发者巨大的便利。

并发不是并行：Concurrency Is Not Parallelism
并发主要由切换时间片来实现“同时”运行，在并行则是直接利用
多核实现多线程的运行，但 Go 可以设置使用核数，以发挥多核计算机
的能力。

Goroutine 奉行通过通信来共享内存，而不是共享内存来通信。

Channel

Channel 是 goroutine 沟通的桥梁，大都是阻塞同步的
通过 make 创建，close 关闭
Channel 是引用类型
可以使用 for range 来迭代不断操作 channel
可以设置单向或双向通道
可以设置缓存大小，在未被填满前不会发生阻塞

Select

可处理一个或多个 channel 的发送与接收
同时有多个可用的 channel时按随机顺序处理
可用空的 select 来阻塞 main 函数
可设置超时

*/

func MyChannel() {
	Println(">>>>>>>>>>>>>>>>>>>>> MyChannel <<<<<<<<<<<<<<<<<<<<<<<<<<<<<")

	channel1()
	channel2()
	channel3()
	channel4()
	channel5()
	// channel6()
	// channel7()
	// channel8()

}

func Gowg(wg *sync.WaitGroup, index int) {
	a := 0
	for i := 0; i < 1000; i++ {
		a += i
	}
	Println("index=", index, ",a=", a)

	wg.Done()
}

func Go(b chan bool, index int) {
	a := 0
	for i := 0; i < 10000; i++ {
		a += i
	}
	Println("index=", index, ",a=", a)

	b <- true

}

func channel1() {
	Println("..................channel1 ..........")
	// 一般channel的声明形式为：	var chanName chan ElementType

	var ch chan int
	// var m map[string] chan bool
	// 声明并初始化了一个int型的名为ch的channel。
	ch = make(chan int)
	// 在channel的用法中，最常见的包括写入和读出。将一个数据写入（发送）
	// 至channel的语法很直观，如下：
	var value int = 2
	go func() {
		ch <- value
		close(ch) //必须有close（ch），与 for：v :=range ch 对应。
		Println("read into channel ch.")
	}()
	// 向channel写入数据通常会导致程序阻塞，直到有其他goroutine从这个
	// channel中读取数据。从	channel中读取数据的语法是
	// value =	<-ch
	for v := range ch {
		Println(v)

	}

	// 如果channel之前没有写入数据，那么从channel中读取数据也会导致程序阻塞，
	// 直到channel	中被写入数据为止。我们之后还会提到如何控制channel只接受
	// 写或者只允许读取，即单向	channel。
}

func channel2() {
	Println("..................channel2 ..........")

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Gowg(&wg, i)
	}
	wg.Wait()
	Println("Gowg end.")

}

func channel3() {
	Println("..................channel3 ..........")
	//-------------------
	// runtime.GOMAXPROCS(runtime.NumCPU())
	b := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go Go(b, i)
	}
	for i := 0; i < 10; i++ {
		<-b
	}

	time.Sleep(2 * time.Microsecond)
	Println("Go end.")
}

func channel4() {
	Println("..................channel4 ..........")
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool)
	a := func() {
		Println("Go Go Go!!!")
		c <- true
		close(c)
	}
	go a()
	for v := range c {
		Println(v)
	}
}

func channel5() {
	Println("..................channel5 ..........")
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case ch <- 0:
		case ch <- 1:
		}
		i := <-ch
		Println("Value received:", i)
	}
}
