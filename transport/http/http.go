package http

import (
	"net/http"
)

// This var block is temporary as we do not want to introduce a global state.
var (
	// DefaultTransport sets a Server transport to HTTP by default, unless configured otherwise.
	DefaultTransport = http.Transport{}
)
