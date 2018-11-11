package packagetest

import (
	"fmt"
	"time"
)
func WhanTimer(){
	fmt.Println(">>>>>>>>>>>>>>>> Whan 延时的3种方法<<<<<<<<<<<<<<<<<<")
	timer1:=time.NewTimer(4*time.Second)
	timer1.Reset(1*time.Second)
	fmt.Println(time.Now())
	
	t:=<-timer1.C
	fmt.Println("延时1秒时间到：",t)
	time.Sleep(1*time.Second)
	fmt.Println("延时1秒时间到")
	
	t=<-time.After(2*time.Second)
	
	fmt.Println("延时2秒时间到",t)
	
//	timer.Stop()
}