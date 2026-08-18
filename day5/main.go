package main

import "fmt"

type user struct {
	id    string
	name  string
	email string
}

func main() {
	var users []user = []user{
		{
			id:    "1",
			name:  "yonis",
			email: "yonis@gmail.com",
		},
		{
			id:    "2",
			name:  "kiino",
			email: "kiino@gmail.com",
		},
		{
			id:    "3",
			name:  "mustafe",
			email: "mustafe@gmail.com",
		},
	}

	fmt.Printf("%#v", users)
}
