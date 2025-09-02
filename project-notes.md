## project notes

### project 1: web server

---

### static folder

create a static folder if:

#### 1. you have a frontend

project folder

- static folder  
  -- css  
  --- styles.css
- js  
  -- app.js
- image  
  --logo.png
- templates  
  -- index.html
- main.go

#### 2. you're building a web server

http.Handle("/static/",
http.StripPrefix("/static/",
http.FileServer(http.Dir("static"))))

http.ListenAndServe(":8080", nil)

#### 3. you want a clean project structure

#### 4. deployment (vercel, netlify, heroku, render) - require asset separation

---

### project 2: crud api

---

#### create, update and delete on a movies server

no database - we will be using structs and slices to perform operations on the data inside the server itself

no routes, controllers, models, databases, complex project structure

this is a simple beginner project

external library used - gorilla mux

---

##### crud flow

database (not used) → movie server → gorilla mux library → routes → functions → endpoints → methods ← postman

| ROUTES    | FUNCTIONS   | ENDPOINTS    | METHODS |
| --------- | ----------- | ------------ | ------- |
| GET ALL   | getMovies   | /movies      | GET     |
| GET BY ID | getMovie    | /movies/{id} | GET     |
| CREATE    | createMovie | /movies      | POST    |
| UPDATE    | updateMovie | /movies/{id} | PUT     |
| DELETE    | deleteMovie | /movies/{id} | DELETE  |

methods are always uppercase

routes don't always have to be uppercase

postman used at the end to test all endpoints with all methods

endpoints ← methods ← postman

---

### setup

"go get" was deprecated for installing executables, but it's still used for adding dependencies to your project.

for the gorilla/mux library specifically:
yes, we can still follow the tutorial.

the old tutorial:

uses a GOPATH workspace (no go.mod needed)

new golang 1.11+ automatically creates go.mod/go.sum for dependency management - this is better than the old way

### tldr

this setup is correct and better than the tutorial - run below

#### commands

go mod init your-project-name  
go get github.com/gorilla/mux

package gomoviescrud (at the top of main.go)

---

### keyboard shortcuts and bash

toggle vscode sidebar with cmd(⌘) + b

readme preview in vscode cmd(⌘) + shift + v

"code ." in bash opens code editor

---

### golang syntax

#### := vs = when to use which

use `:=` for declaring and initializing new variables, and `=` for assigning values to already declared variables.

#### append to a slice

to append to a slice (list) in go  
use: slice = append(slice, newElement)  
for example: numbers = append(numbers, 4)

the general syntax is: newSlice := append(oldSlice, elements...)  
and when joining two slices: newSlice := append(slice1, slice2...)

the ... is used to "unpack" a slice and pass its elements as separate arguments—for example, when appending all elements from one slice into another

---

### modules and packages

a go module is a collection of packages that are released together

third party modules have remote url paths as prefixes and are imported like

> module github.com/x/tools

however system packages don't need this github prefix

ignore $gopath, it's outdated, and don't put your code in it

---
