package services

import (
	"context"
	"io"
	"net"

	"github.com/charoleizer/thuigsinn/ms-users/pkg/proto/authentication"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn *grpc.ClientConn
}

func NewAuthenticationClient(ctx context.Context, addr string) (*Client, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			return (&net.Dialer{}).DialContext(ctx, "tcp", addr)
		}),
	}

	conn, err := grpc.DialContext(context.Background(), addr, opts...)

	conn.Connect()

	if err != nil {
		return nil, err
	}
	return &Client{conn: conn}, nil
}

func (c *Client) AuthenticationRegister(ctx context.Context, req *authentication.RegisterRequest) error {
	client := authentication.NewAuthenticationClient(c.conn)
	stream, err := client.Register(ctx, req)
	if err != nil {
		return err
	}

	for {
		_, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) AuthenticationClose() error {
	return c.conn.Close()
}
