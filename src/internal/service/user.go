package service

import "github.com/manudevelopia/app-go/src/internal/model"

var users = map[int]model.User{
	1: {Name: "John doe", Email: "john.doe@email.com"},
	2: {Name: "Jane doe", Email: "jane.doe@email.com"},
	3: {Name: "Douglas Quake", Email: "douglas.quake@email.com"},
}

func UserAll(email string) []model.User {
	values := []model.User{}
	for _, user := range users {
		if email != "" {
			if user.Email == email {
				values = append(values, user)
			}
		} else {
			values = append(values, user)
		}
	}
	return values
}

func UserById(id int) (model.User, bool) {
	value, exists := users[id]
	return value, exists
}
