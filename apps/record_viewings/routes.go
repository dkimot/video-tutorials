package recordviewings

import "github.com/go-chi/chi"

func (a *app) routes() {
	r := chi.NewRouter()

	r.Post("/{videoID}", a.handleRecordViewing())

	a.router = r
}
