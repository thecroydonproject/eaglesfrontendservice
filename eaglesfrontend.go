package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func main() {

	fmt.Println("ok")

	http.HandleFunc("/results",Showresults)
	http.ListenAndServe(":8080",nil)

}

func Showresults(w http.ResponseWriter, req *http.Request)  {

	t, _ := template.ParseFiles("index.html")

	t.Execute(w,nil)

}