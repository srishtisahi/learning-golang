package routes

import (
	"github.com/gorilla/mux"
	"github.com/srishtisahi/go-mysql/pkg/controllers" 
	// imports the controllers folder and the files inside it
	// golang has absolute paths, you need to write all of this
)

// the routes will help us get the control to the controller
// the main logic is in the controller
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methhods("POST")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
