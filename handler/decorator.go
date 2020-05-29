package handler

import "github.com/mikudos/mikudos_golang_app/interfaces"

// Decorate Decorate
func Decorate(s interfaces.IServer) {
	Handler.Server = s
	s.Configure(Handler)
}
