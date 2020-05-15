package recordviewings

import (
	"net/http"

	"github.com/go-chi/chi"
	goes "github.com/jetbasrawi/go.geteventstore"

	"github.com/dkimot/video-tutorials/pkg/httpserver"
)

func New(esClient goes.Client) http.HandlerFunc {
	a := app{
		Server: httpserver.New(),
		actions: newActions(esClient),
	}
	a.routes()

	return a.router.ServeHTTP
}

type app struct {
	httpserver.Server
	actions *actions
	router  chi.Router
}
