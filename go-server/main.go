package main

import (
	"fmt"
	"log"
	"net/http"
)

// building a simple web server

// we are making a server with three routes

// 1. server -> / -> index.html
// 2. server -> /hello -> hello func
// 3. server -> /form -> form func -> form.html

// net/http is a webserver library
// log is a logging library
// example `2025/08/14 19:45:03 Server started on port 8080`

/*
1. w -> http.ResponseWriter
2. r -> http.Request

3. r.ParseForm() -> parses the form and stores the values in r.Form
4. err != nil

5. fmt.Fprintf

6. r.formValue
*/

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err) // %v is a placeholder for the error
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)       // %s is a placeholder for the string
	fmt.Fprintf(w, "Address = %s\n", address) // %s is a placeholder for the string
	return                                    // don't forget this or you'll get 404 error
}

// remember this syntax, it will be used a lot:

/*
http functions to remember:
1. w -> http.ResponseWriter
2. r -> http.Request

3. r.URL.Path
4. http.Error
5. http.StatusNotFound

6. r.Method != "GET"

7. fmt.Fprintf
*/

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!") // fprintf is a formatted print function
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

/*
main function to remember:
1. http.FileServer
2. http.Dir // directory of the static files

3. http.Handle
4. http.HandleFunc

5. fmt.Printf
6. http.ListenAndServe // starts the server
7. log.Fatal // logs the error and exits the program
*/
