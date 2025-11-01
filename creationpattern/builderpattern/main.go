package main

import (
	"builderpattern/builder"
	"fmt"
	"time"
)

func main() {

	userBuilder := &builder.UserBuilder{}
	userBuilder.SetID("1")
	userBuilder.SetUsername("John Doe")
	userBuilder.SetEmail("john.doe@example.com")
	userBuilder.SetPassword("password")
	userBuilder.SetCreatedAt(time.Now())
	userBuilder.SetUpdatedAt(time.Now())
	user := userBuilder.Build()
	fmt.Println(user.GetID())
	fmt.Println(user.GetUsername())
	fmt.Println(user.GetEmail())
	fmt.Println(user.GetPassword())
	fmt.Println(user.GetCreatedAt())
	fmt.Println(user.GetUpdatedAt())
}
