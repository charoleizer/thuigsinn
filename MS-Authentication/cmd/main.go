package main

import "github.com/charoleizer/thuigsinn/ms-authentication/internal"

func main() {
	db, err := internal.ConnectToDatabase()
	if err != nil {
		panic(err)
	}

	err = internal.RunServer(db, "8081")
	if err != nil {
		panic(err)
	}
}
