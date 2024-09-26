package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nandlabs/golly-samples/rest-server/server-routing/models"
	"github.com/nandlabs/golly-samples/rest-server/server-routing/response"
	"github.com/nandlabs/golly-samples/rest-server/server-routing/store"
	"oss.nandlabs.io/golly/rest/server"
)

func UpdateUser(ctx server.Context) {
	pathId, err := ctx.GetParam("id", server.PathParam)
	if err != nil {
		response.Error(ctx.HttpResWriter(), http.StatusBadRequest, "path params not found")
		return
	}

	var item models.Item
	if err = json.NewDecoder(ctx.GetRequest().Body).Decode(&item); err != nil {
		response.Error(ctx.HttpResWriter(), http.StatusBadRequest, "invalid request payload")
		return
	}

	initStore := store.GetStore()

	if _, exists := initStore.GetById(pathId); !exists {
		response.Error(ctx.HttpResWriter(), http.StatusNotFound, "item not found")
		return
	}

	item.ID = pathId
	initStore.Put(pathId, item)
	response.JSON(ctx.HttpResWriter(), http.StatusOK, item)
}
