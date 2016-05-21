package main

import (
	"fmt"
	"net/http"
	"html/template"
	//"io/ioutil"
	//"encoding/json"
	"io/ioutil"

	"encoding/json"
)



//Game describes results of football game
type Game struct {

Gamedate string `json:"gamedate"`
Team string `json:"team"`
Awayorhome string `json:"awayorhome"`
Competition string `json:"competition"`
Result string `json:"result"`
Score1 string `json:"score1"`
Score2 string `json:"score2"`
}

var Content []Game

func main() {

	fmt.Println("ok")

	http.HandleFunc("/results",Showresults)
	http.ListenAndServe(":8080",nil)

}

func Showresults(w http.ResponseWriter, req *http.Request)  {



	b, _:= http.Get("http://localhost:9000/results")

	a, _ := ioutil.ReadAll(b.Body)

	r := json.Unmarshal(a,&Content)

	if r != nil {}

	fmt.Println(Content)


	t, _ := template.ParseFiles("index.html")

	t.Execute(w,nil)

}