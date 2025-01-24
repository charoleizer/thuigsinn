package internal

import (
	"fmt"
	"os"

	"github.com/nats-io/nats.go"
)

func ConnectToBroker() (nats.JetStreamContext, error) {
	brokerUri := fmt.Sprintf("nats://%s:%s",
		os.Getenv("NATS_HOST"),
		os.Getenv("NATS_PORT"),
	)

	// opts := []nats.Option{
	// 	nats.UserInfo(os.Getenv("NATS_USERNAME"), os.Getenv("NATS_PASSWORD")),
	// }

	nc, err := nats.Connect(brokerUri)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to NATS: ", brokerUri)

	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	streamName := "MS-USERS"
	subjects := []string{"users.>"}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     streamName,
		Subjects: subjects,
		Storage:  nats.FileStorage,
	})
	if err != nil && err != nats.ErrStreamNameAlreadyInUse {
		return nil, err
	}

	return js, nil
}
