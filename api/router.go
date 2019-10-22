package main

import (
	"github.com/gorilla/mux"
	"os"
)

func routesSetup() error {
	router := mux.NewRouter()

	subdir := pathStrip(os.Getenv("ORIGIN"))
	if subdir != "" {
		router = router.PathPrefix(subdir).Subrouter()
	}

	if err := apiRouterInit(router); err != nil {
		return err
	}

	if err := staticRouterInit(router); err != nil {
		return err
	}
	
	return nil
}
