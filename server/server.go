package server

// Server encompasses the top level abstraction of a micro service. It contains a Service, which is responsible for
// your business logic. Server is the basic building block, and creating one will ensure something is up and running,
// even if there are no registered handlers. When running a Server, this will register itself in the service registry.
// By default a Server would use TCP as transport.
type Server interface {
	Run() error
}

