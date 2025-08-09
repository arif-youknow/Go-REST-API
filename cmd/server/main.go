package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
}

//Task type er ekta slice globally create korlam
var myTask []Task


//=========get request=============

func getTasks(w http.ResponseWriter, r *http.Request){

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(myTask)

}

func createTask(w http.ResponseWriter, r *http.Request){


var task Task

json.NewDecoder(r.Body).Decode(&task)
task.ID = len(myTask) + 1
myTask = append(myTask, task)

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(myTask)

}

func registration(w http.ResponseWriter, r *http.Request){


fmt.Fprintln(w, "This is registration page")

}

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


http.HandleFunc("/task" , func(w http.ResponseWriter, r *http.Request) {

	if(r.Method=="GET"){

		getTasks(w, r)
	}else if(r.Method=="POST"){
		createTask(w,r)
	}else{
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not allowed")
	}


})

http.HandleFunc("/", home)
http.HandleFunc("/about", about)
http.HandleFunc("/contact", contact)







fmt.Println("server running: http://localhost:3000/")
http.ListenAndServe(":3000", nil)
}

