package main

import (
	"github.com/fzdwx/iter/fx"
	"github.com/fzdwx/iter/stream"
)

func main() {
	userService := UserService{}

	users := userService.ListUser()
	stream.Of(users).
		Filter(func(u user) bool {
			return u.Age > 18
		}).
		MapTo(func(u user) string {
			return u.Name
		}).
		Distinct(fx.IdentityString).
		ToArray()

}

type user struct {
	Name string
	Age  int
}

type UserService struct {
}

func (u UserService) ListUser() []user {
	return []user{}
}
