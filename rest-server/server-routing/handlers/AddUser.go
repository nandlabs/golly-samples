package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nandlabs/golly-samples/rest-server/server-routing/models"
	"github.com/nandlabs/golly-samples/rest-server/server-routing/response"
	"github.com/nandlabs/golly-samples/rest-server/server-routing/store"
	"oss.nandlabs.io/golly/rest/server"
	"oss.nandlabs.io/golly/uuid"
)

func AddUser(ctx server.Context) {
	var item models.Item
	if err := json.NewDecoder(ctx.GetRequest().Body).Decode(&item); err != nil {
		response.Error(ctx.HttpResWriter(), http.StatusBadRequest, "invalid request payload")
		return
	}

	initStore := store.GetStore()

	uuid, err := uuid.V1()
	if err == nil {
		item.ID = uuid.String()
		initStore.Put(item.ID, item)
		response.JSON(ctx.HttpResWriter(), http.StatusCreated, item)
	} else {
		response.Error(ctx.HttpResWriter(), http.StatusBadRequest, "error generating uuid")
		return
	}
}
