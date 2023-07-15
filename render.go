package main

import(
	"fmt"
	"net/http"
	"html/template"
)


func renderTemplate(w http.ResponseWriter,tmpl string){
	parsedTemplate,_ := template.ParseFiles("templates/" + tmpl)
	fmt.Println(tmpl)
	err := parsedTemplate.Execute(w,nil)
	if err != nil{
		fmt.Println("Error message:",err)
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
		return
	}
}