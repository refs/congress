package http

import (
	"net"
	"net/http"
)

// server implements the Server interface.
type server struct{
	Address string // should be moved to options.
	Handler http.Handler
}

// Run starts a Server.
func (s server) Run() error {
	ln, err := net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}

	return http.Serve(ln, s.Handler)
}

// New returns a new HTTP server.
func New() server {
	return server{}
}
