package transport

import (
	chttp "github.com/refs/congress/pkg/transport/http"
	"net/http"
)

type Transport struct {
	*http.Transport
}

func New() Transport {
	return Transport{
		&chttp.DefaultTransport,
	}
}
