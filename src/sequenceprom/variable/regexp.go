package variable

import (
	"fmt"
	"regexp"

)

func WhanRegexp() {
whanRegexp1()

}

func whanRegexp1() {
	buf := "abc azc a7c aac 888 a9c tac a2c"

	reg1 := regexp.MustCompile(`a.c`)
	if reg1 == nil {
		fmt.Println("err")
		return
	}
	result := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result=", result)
}
