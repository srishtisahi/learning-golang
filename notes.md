## notes
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

### toggle vscode sidebar with cmd + B

---

### CRUD API

#### Create, Update and Delete

