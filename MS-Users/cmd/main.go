package main

import "github.com/charoleizer/thuigsinn/ms-users/internal"

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
