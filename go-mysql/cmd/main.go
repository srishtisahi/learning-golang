package main 

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/srishtisahi/go-mysql/pkg/routes"
)

// create server
// define localhost
// tell go where the routers are

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r) // r here is router in bookstore.go 
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r)) // function to create a server
}