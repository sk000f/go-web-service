package models

import (
	"errors"
	"fmt"
)

// User defines a User type within the web service
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers returns all the user objects
func GetUsers() []*User {
	return users
}

// AddUser creates a new user
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New user must not include ID")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// GetUserByID returns the user with the specified ID
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

// UpdateUser will update the specified user with the details provided
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' could not be found", u.ID)
}

// RemoveUserByID will delete the specified user
func RemoveUserByID(id int) error {
	for i, candidate := range users {
		if candidate.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID '&v' could not be found")
}
