package main

import (
	"github.com/charoleizer/thuigsinn/ms-users/internal"
)

func main() {
	broker, err := internal.ConnectToBroker()
	if err != nil {
		panic(err)
	}

	db, err := internal.ConnectToDatabase()
	if err != nil {
		panic(err)
	}

	infrastructure := internal.Infrastructure{
		Database:   db,
		Broker:     broker,
		ServerPort: "8080",
	}

	err = internal.RunServer(infrastructure)
	if err != nil {
		panic(err)
	}
}
