package main

import (
	"fmt"
	"net/http"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, Gopher!\n")
}

func main() {
	fmt.Println("asssssss")
	// http.HandlerFunc(greeting)这句代码的真正含义是将函数greeting显式转换为HandlerFunc类型
	http.Handle("/",http.HandlerFunc(greeting))
	http.ListenAndServe(":8080", nil)
	//http.ListenAndServe(":8080", http.HandlerFunc(greeting))
	// http.ListenAndServe(":8080", http.HandlerFunc(greeting))
	// http.HandleFunc()
	// mux := http.NewServeMux()
	// http.ListenAndServe(":8080", mux)
}
