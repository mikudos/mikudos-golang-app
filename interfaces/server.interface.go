package interfaces

import (
	pb "github.com/mikudos/mikudos_golang_app/proto/greeter"
)

// IServer Server interface
type IServer interface {
	pb.GreeterServiceServer
}
