package recordviewings

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (a *app) handleRecordViewing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := a.actions.recordViewing(r.Context(), chi.URLParam(r, "videoID"))

		a.ServerError(w, err)
	}
}
