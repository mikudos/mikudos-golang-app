package interfaces

func Decorate(decorated IHandler, s IServer) {
	decorated.ServerDecorater(s)
	s.Configure(decorated)
}
