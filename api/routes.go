package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackzampolin/blockstack-indexer/blockstack"
)

// Route representes an indivdual route in the Blockstack API sever
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a collection of Route
type Routes []Route

// NewRouter returns a router instance to be served
func NewRouter(conf blockstack.ServerConfig) *mux.Router {
	h := NewHandlers(conf)
	routes := Routes{
		Route{
			"V1GetName",
			"GET",
			"/v1/names/{name}",
			h.V1GetNameHandler,
		},
		Route{
			"V1GetNameHistory",
			"GET",
			"/v1/names/{name}/history",
			h.V1GetNameHistoryHandler,
		},
		Route{
			"V1GetNamesInNamespace",
			"GET",
			"/v1/namespaces/{namespace}/names",
			h.V1GetNamesInNamespaceHandler,
		},
		Route{
			"V2GetUserProfile",
			"GET",
			"/v2/users/{name}",
			h.V2GetUserProfileHandler,
		},
		Route{
			"V1GetNameOpsAtHeight",
			"GET",
			"/v1/blockchains/bitcoin/operations/{blockHeight}",
			h.V1GetNameOpsAtHeightHandler,
		},
		Route{
			"V1GetNamesOwnedByAddress",
			"GET",
			"/v1/addresses/bitcoin/{address}",
			h.V1GetNamesOwnedByAddressHandler,
		},
		Route{
			"V1GetZonefile",
			"GET",
			"/v1/names/{name}/zonefile",
			h.V1GetZonefileHandler,
		},
		Route{
			"V1GetNamespaceBlockchainRecord",
			"GET",
			"/v1/namespaces/{namespace}",
			h.V1GetNamespaceBlockchainRecordHandler,
		},
		Route{
			"V1GetNamespaces",
			"GET",
			"/v1/namespace",
			h.V1GetNamespacesHandler,
		},
		Route{
			"V1GetNamesInNamespace",
			"GET",
			"/v1/blockchains/{blockchain}/name_count",
			h.V1GetNamesInNamespaceHandler,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}
