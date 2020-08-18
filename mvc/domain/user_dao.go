package domain

import (
	"fmt"
)

var (
	users = map[int64]*User{
		1: {ID: 1, FirstName: "Rahul", LastName: "Gupta", Email: "guptarahul@gmail.com"},
	}
)

func GetUser(userID int64) (*User, error) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, fmt.Errorf("user with id %v was not found", userID)
}