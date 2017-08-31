package main

import (
	"net/http"

	"github.com/bnkamalesh/webgo"
)

func getRoutes(g *webgo.Globals) []*webgo.Route {
	var mws webgo.Middlewares
	return []*webgo.Route{
		&webgo.Route{
			Name:    "all_options",                       // A label for the API/URI, this is not used anywhere.
			Method:  "OPTIONS",                           // request type
			Pattern: "/:wildcard*",                       // Pattern for the route
			Handler: []http.HandlerFunc{mws.CorsOptions}, // route handler
			G:       g,
		},
		&webgo.Route{
			Name:    "root",                              // A label for the API/URI, this is not used anywhere.
			Method:  "GET",                               // request type
			Pattern: "/",                                 // Pattern for the route
			Handler: []http.HandlerFunc{mws.Cors, dummy}, // route handler
			G:       g,
			FallThroughPostResponse: true,
		},
		&webgo.Route{
			Name:    "auth",                                         // A label for the API/URI, this is not used anywhere.
			Method:  "GET",                                          // request type
			Pattern: "/auth",                                        // Pattern for the route
			Handler: []http.HandlerFunc{mws.Cors, authCheck, dummy}, // route handler
			G:       g,
		},
		&webgo.Route{
			Name:    "mongo",                               // A label for the API/URI, this is not used anywhere.
			Method:  "GET",                                 // request type
			Pattern: "/mgodb/:name",                        // Pattern for the route
			Handler: []http.HandlerFunc{mws.Cors, MongoDB}, // route handler
			G:       g,
		},
		&webgo.Route{
			Name:    "mysql",                             // A label for the API/URI, this is not used anywhere.
			Method:  "GET",                               // request type
			Pattern: "/mysql/:name",                      // Pattern for the route
			Handler: []http.HandlerFunc{mws.Cors, MySQL}, // route handler
			G:       g,
		},
	}
}
