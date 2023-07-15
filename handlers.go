package main

import(
"net/http")





// Home is the home page handler.
func Home(w http.ResponseWriter, r *http.Request){
	renderTemplate(w,"home.page.html")

}

//About - Handles the routes of the about page.
func About (w http.ResponseWriter, r *http.Request){

	renderTemplate(w,"about.page.html")
}
// addValues adds 2 ints and returns the sum


