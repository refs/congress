package transport

import (
	"net/http"
	chttp "github.com/refs/congress/transport/http"
)

type Transport struct {
	*http.Transport
}

func New() Transport {
	return Transport{
		&chttp.DefaultTransport,
	}
}
