package main

import (
	"fmt"
	"net/http"
)


func home(w http.ResponseWriter, r *http.Request){

fmt.Fprintln(w, "This is home page")


}

func contact(w http.ResponseWriter, r *http.Request){

	fmt.Fprintln(w, "This is contact page")
}

func about(w http.ResponseWriter, r *http.Request){

fmt.Fprintln(w, "This is about page")

}


func main(){

http.HandleFunc("/", home)
http.HandleFunc("/about", about)
http.HandleFunc("/contact", contact)







fmt.Println("server running...")
http.ListenAndServe(":3000", nil)
}