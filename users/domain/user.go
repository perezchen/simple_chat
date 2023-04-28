package domain

import (
	"encoding/json"
	"errors"
)

type User struct {
	properties properties
}

type properties struct {
	ID       string `json:"ID"`
	Name     string `json:"Name"`
	LastName string `json:"LastName"`
	Email    string `json:"email"`
}

func (u User) ID() string {
	return u.properties.ID
}
func (u User) Name() string {
	return u.properties.Name
}
func (u User) LastName() string {
	return u.properties.LastName
}
func (u User) Email() string {
	return u.properties.Email
}

func (u User) ToJson() ([]byte, error) {
	body, err := json.Marshal(u.properties)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (u *User) FromJson(aBytes []byte) error {
	err := json.Unmarshal(aBytes, &u.properties)
	if err != nil {
		return err
	}

	return nil
}

func (u User) validate() error {
	if u.properties.ID == "" {
		return errors.New("id is empty")
	}

	return nil
}

func NewUser(id, name, lastname, email string) (User, error) {
	u := User{
		properties: properties{
			ID:       id,
			Name:     name,
			LastName: lastname,
			Email:    email,
		},
	}

	err := u.validate()
	if err != nil {
		return User{}, err
	}

	return u, nil
}
