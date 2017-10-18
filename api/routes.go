package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackzampolin/go-blockstack/blockstack"
	// "github.com/jackzampolin/go-blockstack/indexer"
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

// Routes returns the routes in an array
func (r Routes) Routes() []string {
	var out []string
	for _, rt := range r {
		out = append(out, fmt.Sprintf("%s %s", rt.Method, rt.Pattern))
	}
	return out
}

// NewRouter returns a router instance to be served
func NewRouter(conf blockstack.ServerConfig) *mux.Router {
	h := NewHandlers(conf)
	routes := Routes{
		// // NOTE: Testing Route, Remove
		// Route{
		// 	Name:        "V1GetName",
		// 	Method:      "GET",
		// 	Pattern:     "/resolver/numnames",
		// 	HandlerFunc: h.NumNamesHandler,
		// },
		// // NOTE: Testing Route, Remove
		// Route{
		// 	Name:        "V1GetName",
		// 	Method:      "GET",
		// 	Pattern:     "/resolver/names",
		// 	HandlerFunc: h.GetNamesHandler,
		// },
		Route{
			Name:        "V1GetName",
			Method:      "GET",
			Pattern:     "/v1/names/{name}",
			HandlerFunc: h.V1GetNameHandler,
		},
		Route{
			Name:        "V1GetNameHistory",
			Method:      "GET",
			Pattern:     "/v1/names/{name}/history",
			HandlerFunc: h.V1GetNameHistoryHandler,
		},
		Route{
			Name:        "V1GetNamesInNamespace",
			Method:      "GET",
			Pattern:     "/v1/namespaces/{namespace}/names",
			HandlerFunc: h.V1GetNamesInNamespaceHandler,
		},
		Route{
			Name:        "V2GetUserProfile",
			Method:      "GET",
			Pattern:     "/v2/users/{name}",
			HandlerFunc: h.V2GetUserProfileHandler,
		},
		Route{
			Name:        "V1GetNameOpsAtHeight",
			Method:      "GET",
			Pattern:     "/v1/blockchains/{blockchain}/operations/{blockHeight}",
			HandlerFunc: h.V1GetNameOpsAtHeightHandler,
		},
		Route{
			Name:        "V1GetNamesOwnedByAddress",
			Method:      "GET",
			Pattern:     "/v1/addresses/bitcoin/{address}",
			HandlerFunc: h.V1GetNamesOwnedByAddressHandler,
		},
		Route{
			Name:        "V1GetZonefile",
			Method:      "GET",
			Pattern:     "/v1/names/{name}/zonefile",
			HandlerFunc: h.V1GetZonefileHandler,
		},
		Route{
			Name:        "V1GetNamespaceBlockchainRecord",
			Method:      "GET",
			Pattern:     "/v1/namespaces/{namespace}",
			HandlerFunc: h.V1GetNamespaceBlockchainRecordHandler,
		},
		Route{
			Name:        "V1GetNamespaces",
			Method:      "GET",
			Pattern:     "/v1/namespace",
			HandlerFunc: h.V1GetNamespacesHandler,
		},
		Route{
			Name:        "V1GetNamesInNamespace",
			Method:      "GET",
			Pattern:     "/v1/blockchains/{blockchain}/name_count",
			HandlerFunc: h.V1GetNamesInNamespaceHandler,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}
