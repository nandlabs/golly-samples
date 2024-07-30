package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nandlabs/golly-samples/turbo/server-routing/models"
	"github.com/nandlabs/golly-samples/turbo/server-routing/response"
	"github.com/nandlabs/golly-samples/turbo/server-routing/store"
	"oss.nandlabs.io/golly/turbo"
)

func UpdateUser(router *turbo.Router, store *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathId, err := router.GetPathParams("id", r)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "path params not found")
			return
		}

		var item models.Item
		if err = json.NewDecoder(r.Body).Decode(&item); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		if _, exists := store.GetById(pathId); !exists {
			response.Error(w, http.StatusNotFound, "item not found")
			return
		}

		item.ID = pathId
		store.Put(pathId, item)
		response.JSON(w, http.StatusOK, item)

	}
}
