### bookstore management

#### apis

1. database - mysql
2. gorm (go object relational mapper) - serves as a bridge between the go application and mysql database

   > crud operations

   > auto migrations - update db tables based on go struct definitions

   > database connection pooling

   > orm - supports associations like (has one, has many, belongs to, many to many), polymorphism, and single-table inheritance

3. json marshalling and unmarshalling

   > marshalling - convert go value to json

   > unmarshalling - convert json to go value

   > here, go value can be a struct, map, slice, etc.

4. project structure

5. gorilla mux - http request router (multiplexer)

   > route matching

   > url variables - eg. /books/{id:[0-9]+}

   > sub routers - shared attributes like /api or common middleware only apply to a subset of routes to improve performance

   > url building (reverse routing) - generate links from route names

   > middleware support - because every router implements http.Handler, you can wrap it with logging, CORS, authentication, and other middleware layers

---

folders - cmd and pkg

1. cmd - main.go

2. pkg has more folders - config, controllers, models, routes, utils

   > config - app.go

   > controllers - book-controller

   > models - book.go

   > routes - bookstore-routes

   > utils - utils.go

3. routes

   > controller functions - get books, create book, get book by id, update book, delete book

get books ← /book/ ← get

create book ← /book/ ← post  
uses json marshalling/unmarshalling

get book by id ← /book/{bookid} ← get  
create function to get a book by id

update book ← /book/{bookid} ← put

delete book ← /book/{bookid} ← delete

---

#### packages and modules

installed beforehand from github

> go mod init github.com/srishtisahi/go-mysql  
> go get "github.com/jinzhu/gorm"  
> go get "github.com/jinzhu/gorm/dialects/mysql"  
> go get "github.com/gorilla/mux"

---
