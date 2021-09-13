package endpoints

import (
	"../controllers"
	"../models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func usersListHandler(w http.ResponseWriter, request *http.Request) {
	users := controllers.GetAllUsers()
	data, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func usersGetHandler(w http.ResponseWriter, request *http.Request) {
	userIdParam := strings.Trim(request.URL.Path, "/")
	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Bad user id: "+userIdParam)
		return
	}
	user, ok := controllers.GetUser(uint8(userId))
	if !ok {
		w.WriteHeader(404)
		fmt.Fprintf(w, "User with id %d not found", userId)
		return
	}
	data, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func usersUpdateHandler(w http.ResponseWriter, request *http.Request) {
	userIdParam := strings.Trim(request.URL.Path, "/")
	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Bad user id: "+userIdParam)
		return
	}
	userToUpdate := models.User{}
	err2 := json.NewDecoder(request.Body).Decode(&userToUpdate)
	if err2 != nil || userToUpdate.Name == "" {
		w.WriteHeader(400)
		fmt.Fprint(w, "Bad user data")
		return
	}
	userToUpdate.ID = uint8(userId)
	updatedUser, ok := controllers.UpdateUser(userToUpdate)
	if !ok {
		w.WriteHeader(404)
		fmt.Fprintf(w, "User with id %d not found", userId)
		return
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(updatedUser)
	w.Write(data)
}

func usersCreateHandler(w http.ResponseWriter, request *http.Request) {
	newUser := models.User{}
	err := json.NewDecoder(request.Body).Decode(&newUser)
	if err != nil || newUser.Name == "" {
		w.WriteHeader(400)
		fmt.Fprint(w, "Bad user data")
		return
	}
	newUser = controllers.CreateUser(newUser)
	w.WriteHeader(201)
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(newUser)
	w.Write(data)
}

func usersDeleteHandler(w http.ResponseWriter, request *http.Request) {
	userIdParam := strings.Trim(request.URL.Path, "/")
	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Bad user id: "+userIdParam)
		return
	}
	ok := controllers.DeleteUser(uint8(userId))
	if !ok {
		w.WriteHeader(404)
		fmt.Fprintf(w, "User with id %d not found", userId)
		return
	}
	w.WriteHeader(204)
	w.Header().Set("Content-Type", "application/json")
}

func UsersEndpoint(w http.ResponseWriter, request *http.Request) {
	fmt.Printf("Got %s request at %s \n", request.Method, request.URL.Path)
	switch request.Method {
	case http.MethodGet:
		if len(request.URL.Path) > 0 {
			usersGetHandler(w, request)
		} else {
			usersListHandler(w, request)
		}
		break
	case http.MethodPost:
		usersCreateHandler(w, request)
		break
	case http.MethodPut:
		usersUpdateHandler(w, request)
		break
	case http.MethodDelete:
		usersDeleteHandler(w, request)
	default:
		w.WriteHeader(400)
		fmt.Fprint(w, "Bad request")
	}
}
