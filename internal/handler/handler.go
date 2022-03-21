package handler

import (
	"encoding/json"
	"first_project/internal/user"
	"fmt"
	"log"
	"net/http"
)

type createUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  int    `json:"sex"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var body createUserRequest
	json.NewDecoder(r.Body).Decode(&body)
	err := user.NewUser(body.Name, body.Age, body.Sex)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	log.Printf("Create new user: %v", body)
	fmt.Fprintf(w, "OK")
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	users := user.GetUser()
	log.Println("Give all users")
	fmt.Fprintf(w, "%v", users)
}

type getUserbyIdRequest struct {
	Id int `json:"id"`
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	var req getUserbyIdRequest
	json.NewDecoder(r.Body).Decode(&req)
	user, err := user.GetUserById(req.Id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
		return

	}
	log.Printf("Give user: %v", user)
	fmt.Fprintf(w, "%+v", user)
}

type DeleteRequest struct {
	Id int `json:"id"`
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var req DeleteRequest
	json.NewDecoder(r.Body).Decode(&req)
	user.DeleteUser(req.Id)
}

func Pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong")
}
