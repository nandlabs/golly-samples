package handlers

import (
	"net/http"

	"github.com/nandlabs/golly-samples/turbo/server-routing/response"
	"github.com/nandlabs/golly-samples/turbo/server-routing/store"
	"oss.nandlabs.io/golly/turbo"
)

func DeleteUser(router *turbo.Router, store *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathId, err := router.GetPathParams("id", r)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "path params not found")
			return
		}

		if _, exists := store.GetById(pathId); !exists {
			response.Error(w, http.StatusNotFound, "item not found")
			return
		}

		store.Delete(pathId)
		response.JSON(w, http.StatusOK, map[string]string{"result": "success"})

	}
}
