package interfaces

import (
	pb "github.com/mikudos/mikudos_golang_app/proto/greeter"
)

// IServer Server interface
type IServer interface {
	pb.GreeterServiceServer
	Configure(*IHandler)
	SetDB(interface{}) error
	GetDB(dbType string) (interface{}, error)
}
