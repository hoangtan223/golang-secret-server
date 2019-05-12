package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var store = newInMemoryStore()

var routes = Routes{
	Route{
		"SecretCreate",
		"POST",
		"/secret",
		SecretCreate(store),
	},
	Route{
		"SecretGet",
		"GET",
		"/secret/{secretHash}",
		SecretGet,
	},
}
