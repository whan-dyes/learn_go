package goweb

import (
	// "bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	// "os"
)

func goweb5() {
	fmt.Println("----------------goweb5------------")
}

func goweb4() {
	fmt.Println("----------------goweb4------------")
}

func goweb3() {
	fmt.Println("----------------goweb3------------")
}

func goweb2() {
	fmt.Println("----------------goweb2------------")
}

func goweb1() {
	fmt.Println("----------------goweb1------------")
	http.HandleFunc("/", index)
	http.HandleFunc("/doLogin", doLogin)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/logout", logoutHandler)
	err := http.ListenAndServe("localhost:8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}

func index(rw http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("../goweb/index.tpl")

	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
	t.Execute(rw, nil)
	// fmt.Fprint(rw, "index out!")
}

func doLogin(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Fprintf(rw, "用户名:", req.Form["username"][0], "密码:", req.Form["password"][0])
}

func logoutHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "Log out!")
}

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if len(req.Form["name"]) > 0 {
		fmt.Fprint(rw, "Hello ", req.Form["name"][0])
	} else {
		fmt.Fprint(rw, "Hello world!")
	}

}
func MyGoweb() {
	fmt.Println(">>>>>>>>>>>>>>>>>> MyGoweb <<<<<<<<<<<<<<<<<<<<<")
	goweb1()
	goweb2()
	goweb3()
	goweb4()
	goweb5()
}
