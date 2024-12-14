package internal

import (
	"time"

	"github.com/charoleizer/thuigsinn/ms-authentication/internal/handlers"
	"github.com/charoleizer/thuigsinn/ms-authentication/pkg/proto/authentication"
	"google.golang.org/grpc"
)

type server struct {
	authentication.UnimplementedAuthenticationServer
	handlers *handlers.Handlers
}

func RegisterHandlers(s *grpc.Server, h *handlers.Handlers) {
	authentication.RegisterAuthenticationServer(s, &server{handlers: h})
}

func (s *server) Register(req *authentication.RegisterRequest, stream authentication.Authentication_RegisterServer) error {
	responseChan := make(chan *authentication.RegisterResponse)

	go func() {
		response := &authentication.RegisterResponse{
			Status: authentication.Status_Received,
		}
		err := s.handlers.Register(req, response)
		if err != nil {
			response.Status = authentication.Status_Failed
		}
		responseChan <- response
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case response := <-responseChan:
			if err := stream.Send(response); err != nil {
				return err
			}
			if response.Status == authentication.Status_Completed || response.Status == authentication.Status_Failed {
				return nil
			}
		case <-ticker.C:
			response := &authentication.RegisterResponse{
				Status: authentication.Status_Processing,
			}
			if err := stream.Send(response); err != nil {
				return err
			}
		}
	}

}
