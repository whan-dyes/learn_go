package other

import (
	"fmt"
	"os"
	"reflect"
)

func init() {
	fmt.Println("This is a other/Other() init func1")
}

func Other() {
	fmt.Println(">>>>>>>>>>>>>>>>>> other <<<<<<<<<<<<<<<<<<<<<<<")
	other1()
}

func other1() {
	fmt.Println("-----------------other1-----------------")
	args := os.Args
	len := len(args)
	fmt.Printf("os.args length=%d\n", len)
	for i := 0; i < len; i++ {
		fmt.Printf("第 %d 个参数为%v ", i, args[i])
		fmt.Println("类型为：", reflect.TypeOf(args[i]).String())
	}

	fmt.Println()

	for i, v := range args {
		fmt.Printf("第 %d 个参数为%v ,类型为%T\n", i, v, v)
	}

}
