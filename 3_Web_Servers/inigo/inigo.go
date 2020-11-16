package main

import (
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, req*http.Request){
	fmt.Fprint(res, "Hello my name is Shayan Riyaz")
}

func main(){
	http.HandleFunc("/",hello)
	http.ListenAndServe("localhost:4000",nil)
}
