package api

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

type APIConfFunc func(*API)

type API struct {
	addr   string
	router *mux.Router
	cors   *cors.Cors
}

func New(addr string, origins, methods []string) *API {
	apisrv := &API{
		addr:   addr,
		router: mux.NewRouter(),
		cors: cors.New(cors.Options{
			AllowedOrigins: origins,
			AllowedMethods: methods,
		}),
	}

	return apisrv
}

// Listen starts the HTTP Server and listens on the given
// host address
func (api *API) Listen() error {
	listener, err := net.Listen("tcp", api.addr)
	if err != nil {
		return errors.Wrap(err, "creating net.Listener")
	}
	logrus.Infof("[api.Listen] listening on %q\n", api.addr)

	return http.Serve(listener, cors.Default().Handler(api.router))
}

func (api *API) Register(path string, f http.HandlerFunc, method string) {
	logrus.Infof("[api.Register] registered route: %s\n", path)
	api.router.HandleFunc(path, f).Methods(method)
}
