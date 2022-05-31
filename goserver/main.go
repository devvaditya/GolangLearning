package main

import (
	"fmt"
	"log"
	"net/http"
)



// form handler
// Post request
func funcFormHandler (w http.ResponseWriter,r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
		 
	
	// else

}





// hello handler 
// w ->
// request user send
func funcHelloHandler(w http.ResponseWriter,r *http.Request){
	// if path is wrong or something error
	
	if r.URL.Path != "/hello"{
		http.Error(w,"404 not found", http.StatusNotFound)
		return

	}
	// if method is post 
	if r.Method != "GET" {
		http.Error(w,"method is not supported",http.StatusNotFound)
		return 
	}
	fmt.Fprintf(w,"hello")

}


func main()  {
	// we are telling go to check the static file and it will serve us index.html file 
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	// handle func is coming from http package
	http.HandleFunc("/form",funcFormHandler)
	http.HandleFunc("/hello",funcHelloHandler)

	// server start port
	fmt.Printf("server start at port 8080")
	// handle error either set the port at 8080 else return nil
	if err := http.ListenAndServe(":8080",nil); err != nil{
		// here we are using log package
		log.Fatal(err);
	}


}