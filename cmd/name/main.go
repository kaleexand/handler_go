package main

import (
	"first_project/internal/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", handler.Pong)
	http.HandleFunc("/create", handler.CreateUser)
	http.HandleFunc("/users", handler.GetAllUser)
	http.HandleFunc("/user", handler.GetUserById)
	http.HandleFunc("/delete", handler.DeleteUser)
	http.ListenAndServe(":8080", nil)
}
