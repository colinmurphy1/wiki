package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/colinmurphy1/wiki/routes"
	"github.com/colinmurphy1/wiki/state"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Initialize list of users and configuration
	err := state.Init("./exampleSite/wiki.yaml")
	if err != nil {
		log.Fatalf("Error loading program: %s\n", err)
		return
	}

	// Create a new router and add routes
	rtr := httprouter.New()
	rtr.GET("/", routes.RenderPage)
	rtr.GET("/p/*page", routes.RenderPage) // wildcard route

	// Login
	rtr.GET("/login", routes.LoginForm)
	rtr.POST("/login", routes.Login)

	// Start http server
	err = http.ListenAndServe(fmt.Sprintf("%s:%d", state.Conf.Server.Address, state.Conf.Server.Port), rtr)
	if err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
	//log.Printf("Server up and listening on %s:%d", state.Conf.Server.Address, state.Conf.Server.Port)
}
