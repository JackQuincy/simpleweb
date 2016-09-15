package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
	var count int
	count = 1
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) { 
		w.Header().Set("Content-Type", "text/html; charset=utf-8") 
		fmt.Fprintln(w, "<h1>Hello</h1>") 
		fmt.Fprintln(w, "<br>This server has served", count, " requests")
		count++ 
	}) 
	log.Fatal(http.ListenAndServe(":80", nil))

}

