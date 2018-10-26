package channel

import (
	"fmt"
	"runtime"
	"sync"
)

func MyChannel() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool)
	a := func() {
		fmt.Println("Go Go Go!!!")
		c <- true
		close(c)
	}
	go a()
	for v := range c {
		fmt.Println(v)
	}
	//-------------------
	b := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go(b, i)
	}
	for i := 0; i < 10; i++ {
		<-b

	}
	fmt.Println("Go end.")
	//-----------------------
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Gowg(&wg, i)
	}
	wg.Wait()
	fmt.Println("Gowg end.")

}

func Go(b chan bool, index int) {
	a := 0
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println("index=", index, ",a=", a)

	b <- true

}
func Gowg(wg *sync.WaitGroup, index int) {
	a := 0
	for i := 0; i < 1000; i++ {
		a += i
	}
	fmt.Println("index=", index, ",a=", a)

	wg.Done()
}
