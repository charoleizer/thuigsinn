package main

import (
	"users/internal"
)

func main() {
	db, err := internal.ConnectToDatabase()
	if err != nil {
		panic(err)
	}

	err = internal.RunServer(db, "8080")
	if err != nil {
		panic(err)
	}
}
