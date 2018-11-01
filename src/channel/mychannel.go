package channel

import (
	. "fmt"
	"runtime"
	"time"
	"sync"
)

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
		close(ch)    //必须有close（ch），与 for：v :=range ch 对应。
		Println("read into channel ch.")
	}()
	// 向channel写入数据通常会导致程序阻塞，直到有其他goroutine从这个
	// channel中读取数据。从	channel中读取数据的语法是
// value =	<-ch
for v :=range ch{
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

	for 	i := 0; i < 10; i++ {
		go Go(b, i)
	}
	for i := 0; i < 10; i++ {
		<-b
	}

	time.Sleep(2*time.Microsecond)
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
	for i:=0;i<10;i++ {
		select {
			case ch <- 0:
			case ch <- 1:
		}
	i := <-ch
	Println("Value received:", i)
	}
}
