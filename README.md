# Go-MVC

This is my personal project to learn Go and MVC architecture.
Below is the project structure.

```shell

mywebapp/
|-- controllers/
|   |-- home_controller.go
|   |-- user_controller.go
|-- models/
|   |-- user.go
|-- views/
|   |-- home/
|   |   |-- index.html
|   |-- user/
|   |   |-- index.html
|-- middleware/
|   |-- middleware.go
|-- helper/
|   |-- helper.go
|-- config/
|   |-- config.go
|   |-- database.go
|-- routes/
|   |-- api.go
|   |-- web.go
|-- storage/
|   |-- uploads/
|   |   |-- product-001.png
|   |-- logs/
|   |   |-- 2023-05-01.log
|-- public/
|   |-- images/
|   |   |-- favicon.png
|   |-- css/
|   |   |-- styles.css
|   |-- js/
|   |   |-- script.js
|   |-- fonts/
|       |-- font.ttf
|-- main.go
|-- go.mod
|-- go.sum
|-- .env

```

## 1. Routes

I will use Gin framework to handle routes.
Routes are defined in `routes/web.go` and `routes/api.go` files.

