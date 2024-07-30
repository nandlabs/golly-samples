package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nandlabs/golly-samples/turbo/server-routing/models"
	"github.com/nandlabs/golly-samples/turbo/server-routing/response"
	"github.com/nandlabs/golly-samples/turbo/server-routing/store"
	"oss.nandlabs.io/golly/uuid"
)

func AddUser(store *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item models.Item
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		uuid, err := uuid.V1()
		if err == nil {
			item.ID = uuid.String()
			store.Put(item.ID, item)
			response.JSON(w, http.StatusCreated, item)
		} else {
			response.Error(w, http.StatusBadRequest, "error generating uuid")
			return
		}
	}
}
