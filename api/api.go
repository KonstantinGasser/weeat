package api

import (
	"net"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type API struct {
	addr string
}

func New(addr string) *API {
	return &API{
		addr: addr,
	}
}

// Listen starts the HTTP Server and listens on the given
// host address
func (api *API) Listen() error {
	listener, err := net.Listen("tcp", api.addr)
	if err != nil {
		return errors.Wrap(err, "creating net.Listener")
	}
	logrus.Infof("[api.Listen] listening on %q\n", api.addr)
	return http.Serve(listener, nil)
}

func (api *API) Register(path string, method methodIntercepter, f http.HandlerFunc) {
	logrus.Infof("[api.Register] registered route: %s\n", path)
	http.HandleFunc(path, method(f))
}
