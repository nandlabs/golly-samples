package handlers

import (
	"net/http"

	"github.com/nandlabs/golly-samples/rest-server/server-routing/response"
	"github.com/nandlabs/golly-samples/rest-server/server-routing/store"
	"oss.nandlabs.io/golly/rest"
)

func DeleteUser(ctx rest.ServerContext) {
	pathId, err := ctx.GetParam("id", rest.PathParam)
	if err != nil {
		response.Error(ctx.HttpResWriter(), http.StatusBadRequest, "path params not found")
		return
	}

	initStore := store.GetStore()

	if _, exists := initStore.GetById(pathId); !exists {
		response.Error(ctx.HttpResWriter(), http.StatusNotFound, "item not found")
		return
	}

	initStore.Delete(pathId)
	response.JSON(ctx.HttpResWriter(), http.StatusOK, map[string]string{"result": "success"})
}
