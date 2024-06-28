package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	userId, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("Iserted id: %d", userId)))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Users"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
