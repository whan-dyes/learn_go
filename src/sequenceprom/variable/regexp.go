package variable

import (
	"fmt"
	"regexp"

)

func WhanRegexp() {
fmt.Println(">>>>>>>>>>>> WhanRegexp <<<<<<<<<<<<<<<<<<<")
whanRegexp1()

}

func whanRegexp1() {
	fmt.Println("...................1")
	buf := "abc azc a7c aac 888 a9c tac a2c"

	reg1 := regexp.MustCompile(`a.c`)
	if reg1 == nil {
		fmt.Println("err")
		return
	}
	result := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result=", result)
		fmt.Println("...................2")
	buf = "3.14 567 agsdg 1.23.7. 8.9 lsdljgl 6.66 7.8"

	reg1 = regexp.MustCompile(`\d+\.\d+`)
	if reg1 == nil {
		fmt.Println("err")
		return
	}
	result = reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result=", result)
}
