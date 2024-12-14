package errdefs

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Validations errors
var (
// ErrUsernameRequired = status.Error(codes.InvalidArgument, "username is required")
// ErrEmailRequired    = status.Error(codes.InvalidArgument, "email is required")
// ErrPasswordRequired = status.Error(codes.InvalidArgument, "password is required")

// ErrEmailAlreadyExists    = status.Error(codes.AlreadyExists, "email already exists")
// ErrUsernameAlreadyExists = status.Error(codes.AlreadyExists, "username already exists")

//	ErrInvalidUsername = func(message string) error {
//		return status.Error(codes.InvalidArgument, "invalid username: "+message)
//	}
//
//	ErrInvalidEmail = func(message string) error {
//		return status.Error(codes.InvalidArgument, "invalid email: "+message)
//	}
//
//	ErrInvalidPassword = func(message string) error {
//		return status.Error(codes.InvalidArgument, "invalid password: "+message)
//	}
)

// Database errors
var (
	// ErrUserNotFound = status.Error(codes.NotFound, "user not found")

	//	ErrDatabaseCountDocuments = func(message string) error {
	//		return status.Error(codes.Internal, "failed to count documents: "+message)
	//	}
	//
	ErrDatabaseInsertOne = func(message string) error {
		return status.Error(codes.Internal, "failed to insert one document: "+message)
	}

	ErrDatabaseTimeout = func(message string) error {
		return status.Error(codes.DeadlineExceeded, "database timeout: "+message)
	}

//	ErrDatabaseFindOne = func(message string) error {
//		return status.Error(codes.Internal, "failed to find one document: "+message)
//	}
//
//	ErrDatabaseUpdateOne = func(message string) error {
//		return status.Error(codes.Internal, "failed to update one document: "+message)
//	}
//
//	ErrDatabaseDeleteOne = func(message string) error {
//		return status.Error(codes.Internal, "failed to delete one document: "+message)
//	}
)

// Other errors
var (
	//	ErrHashPassword = func(message string) error {
	//		return status.Error(codes.Internal, "failed to hash password: "+message)
	//	}
	ErrInvalidObjectID = func(message string) error {
		return status.Error(codes.InvalidArgument, "invalid object ID: "+message)
	}
)
