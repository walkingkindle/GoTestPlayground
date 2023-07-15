package main

import (
	"fmt"
	"net/http"

)
const portNumber = "localhost:8080"


func main(){
     
	http.HandleFunc("/",Home)
	http.HandleFunc("/about",About)
	fmt.Println(fmt.Sprintf("Starting application on port %s",portNumber))
	_= http.ListenAndServe(portNumber , nil)

}