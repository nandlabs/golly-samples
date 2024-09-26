package handlers

import (
	"net/http"

	"github.com/nandlabs/golly-samples/rest-server/server-routing/response"
	"github.com/nandlabs/golly-samples/rest-server/server-routing/store"
	"oss.nandlabs.io/golly/rest/server"
)

func GetUsers(ctx server.Context) {
	initStore := store.GetStore()
	items := initStore.GetAll()
	response.JSON(ctx.HttpResWriter(), http.StatusOK, items)
}
