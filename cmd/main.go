package main

import (
	"fmt"

	"github.com/perezchen/simple_chat/users/domain"
)

func main() {
	throwError := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	// Create user with factory
	// user, err := domain.NewUser("hasjfa", "Chencho", "Perez", "chencho@gmail.com")
	// throwError(err)

	// // Encode & Print user to json
	// jsonb, err := user1.ToJson()
	// throwError(err)

	// fmt.Println(string(jsonb))
	userJson := `{"ID":"hasjfa","Name":"Chencho","LastName":"Perez","email":"chencho@gmail.com"}`

	user := domain.User{}

	err := user.FromJson([]byte(userJson))
	throwError(err)

	fmt.Println("user id:", user.ID())
	fmt.Println("user name:", user.Name())
	fmt.Println("user last name:", user.LastName())
	fmt.Println("user email:", user.Email())

}
