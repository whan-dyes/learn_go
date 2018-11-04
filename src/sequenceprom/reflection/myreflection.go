package reflection

import (
	"fmt"
	"reflect"
)

/*
反射reflection

  反射可大大提高程序的灵活性，使得 interface{} 有更大的发挥余地
  反射使用 TypeOf 和 ValueOf 函数从接口中获取目标对象信息
  反射会将匿名字段作为独立字段（匿名字段本质）
  想要利用反射修改对象状态，前提是 interface.data 是 settable，
  即 pointer-interface
  - 通过反射可以“动态”调用方法
*/
type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) HelloName(name string) {
	fmt.Println("Hello", name, " my name is", u.Name)
}

//----------------------------------------------
type Manager struct {
	User
	title string
}

func (u User) Hello() {
	fmt.Println("Hello world.")
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("\n error argument,return ....")
		return
	}

	fmt.Println("Type:", t.Name())
	v := reflect.ValueOf(o)
	fmt.Println("Fields:")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}
func MyReflection() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>> MyReflection <<<<<<<<<<<<<<<<<<<<<<<")
	reflection1()
	reflection2()
	reflection3()
	reflection4()
	reflection5()
	reflection6()
}

func reflection1() {
	fmt.Println("---------- 1 -------------")
	u := User{1, "OK", 12}
	Info(u)

	v := &User{1, "OK", 12}
	Info(v) //不能传入一个指针
}

func reflection2() {
	fmt.Println("---------- 2 -------------")
	m := Manager{User: User{1, "OK", 12}, title: "123"}
	t := reflect.TypeOf(m)
	j := t.NumField()

	for i := 0; i < j; i++ {
		fmt.Printf("%#v\n", t.Field(i))
	}

	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))

}
func reflection3() {
	fmt.Println("---------- 3 -------------")
	x := 123
	fmt.Println(x)
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)

	fmt.Println(x)
	//----------------------------------
	u := User{1, "OK", 12}

	a := func(o interface{}) {
		v := reflect.ValueOf(o)
		fmt.Println("准备设置 User.Name :", v.Elem())
		if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
			fmt.Println("error return...")
			return
		} else {
			v = v.Elem()
		}

		f := v.FieldByName("Name")
		if !f.IsValid() {
			fmt.Println("设置失败。")
			return
		}
		if f.Kind() == reflect.String {
			f.SetString("ByeBye!")
		} else {
			fmt.Println("set string false")
		}
		fmt.Println("设置成功 User.Name :", v)
	}
	a(&u)

}
func reflection4() { //通过反射调用方法： HelloName
	fmt.Println("---------- 4 -------------")
	u := User{1, "OK", 12}
	v := reflect.ValueOf(u)
	mv := v.MethodByName("HelloName")

	args := []reflect.Value{reflect.ValueOf("Joe")}
	mv.Call(args)
}

func reflection5() {
	fmt.Println("---------- 5 -------------")
}
func reflection6() {
	fmt.Println("---------- 6 -------------")
}
