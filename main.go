package main

import (
	"fmt"
	"log"
	"net/http"
)


func main(){
fileServer := http.FileServer(http.Dir("./static"))
http.Handle("/", fileServer)
http.HandleFunc("/hello", helloHandler)
http.HandleFunc("/form", formHandler)

fmt.Printf("starting server at :8080")
if err:= http.ListenAndServe(":8080", nil); err != nil{
	log.Fatal(err)
}
}
func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "parseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w,"POST request successfull")
	name:= r.FormValue("name")
	address:= r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "address= %s\n", address)
}