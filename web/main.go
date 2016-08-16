package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println("begin web.....")
	http.HandleFunc("/api", apiHandler)
	http.ListenAndServe(":8888", nil)
}

func apiHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	m := r.Method
	if m == "POST" {
		fmt.Println("GET post request")
	}
	t, err := template.ParseFiles("template/html/api/index.html")
	if err != nil {
		log.Println(err)
	}
	rw.WriteHeader(5)
	t.Execute(rw, nil)
}
