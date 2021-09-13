package controllers

import (
	"../models"
)

var database map[uint8]models.User
var nextKey uint8 = 1

func ConnectDB() {
	ConnectTestDB()
}

func ConnectTestDB() {
	database = make(map[uint8]models.User)
	var i uint8
	for i = 0; i <= 10; i++ {
		database[nextKey] = models.User{
			ID:   nextKey,
			Name: "Adam",
			Age:  22,
		}
		nextKey++
	}
}

func GetAllUsers() []models.User {
	users := make([]models.User, 0, len(database))
	for _, value := range database {
		users = append(users, value)
	}
	return users
}

func GetUser(userID uint8) (models.User, bool) {
	user, found := database[userID]
	return user, found
}

func CreateUser(user models.User) models.User {
	user.ID = nextKey
	database[user.ID] = user
	nextKey++
	return user
}
func UpdateUser(user models.User) (models.User, bool) {
	currentUser, found := GetUser(user.ID)
	if !found {
		return currentUser, false
	}
	database[user.ID] = user
	return user, true
}

func DeleteUser(userID uint8) bool {
	_, found := database[userID]
	if found {
		delete(database, userID)
		return true
	} else {
		return false
	}
}
