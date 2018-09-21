package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const form = `<html><body><form action="#" method="post" name="bar">
                    <input type="text" name="in"/>
                    <input type="text" name="in"/>
                     <input type="submit" value="Submit"/>
             </form></html></body>`

func SimpleServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>hello, world</h1>")
}

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		request.ParseForm()
		io.WriteString(w, request.Form["in"][1])
		io.WriteString(w, "\n")
		io.WriteString(w, request.FormValue("in"))
	}
}
func main() {

	http.HandleFunc("/test1", logPanics(SimpleServer))
	http.HandleFunc("/test2", logPanics(FormServer))
	if err := http.ListenAndServe("0.0.0.0:8088", nil); err != nil {
	}

	//panic1(ui)(st{"1"})//logPanics封装解释
}

func logPanics(handle http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		handle(writer, request)
	}
}

type st struct {
	a string
}

func ui(x st) {
	fmt.Println(x)
}
func panic1(a func(st)) func(st) {
	return func(x st) {
		a(x)
	}
}
