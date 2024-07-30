package handlers

import (
	"net/http"

	"github.com/nandlabs/golly-samples/turbo/server-routing/response"
	"github.com/nandlabs/golly-samples/turbo/server-routing/store"
)

func GetUsers(store *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := store.GetAll()
		response.JSON(w, http.StatusOK, items)
	}
}
