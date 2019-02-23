package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	hdl "logika/idm/handlers"
)

func main() {
	// defer database.Disconnect()

	router := mux.NewRouter()
	//user route
	router.HandleFunc("/api/idm/users", hdl.GetAllUserEndpoint).Methods("GET")
	router.HandleFunc("/api/idm/user/{id}", hdl.GetUserEndpoint).Methods("GET")
	router.HandleFunc("/api/idm/users/{offset}/{limit}", hdl.GetPagingUserEndpoint).Methods("GET")
	router.HandleFunc("/api/idm/users/{offset}/{limit}/{text}", hdl.SearchUserEndpoint).Methods("GET")
	router.HandleFunc("/api/idm/user", hdl.CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/api/idm/user/{id}", hdl.UpdateUserEndpoint).Methods("PUT")

	//org route
	router.HandleFunc("/api/idm/orgs", hdl.GetAllOrgEndpoint).Methods("GET")
	router.HandleFunc("/api/idm/org/{id}", hdl.GetOrgEndpoint).Methods("GET")
	router.HandleFunc("/api/idm/orgs/{offset}/{limit}", hdl.GetPagingOrgEndpoint).Methods("GET")
	router.HandleFunc("/api/idm/orgs/{offset}/{limit}/{text}", hdl.SearchOrgEndpoint).Methods("GET")
	router.HandleFunc("/api/idm/org", hdl.CreateOrgEndpoint).Methods("POST")
	router.HandleFunc("/api/idm/org/{id}", hdl.UpdateOrgEndpoint).Methods("PUT")

	//org route
	router.HandleFunc("/api/idm/resources", hdl.GetAllResourceEndpoint).Methods("GET")
	router.HandleFunc("/api/idm/resource/{id}", hdl.CreateResourceEndpoint).Methods("GET")
	router.HandleFunc("/api/idm/resources/{offset}/{limit}", hdl.GetPagingResourceEndpoint).Methods("GET")
	router.HandleFunc("/api/idm/resources/{offset}/{limit}/{text}", hdl.SearchResourceEndpoint).Methods("GET")
	router.HandleFunc("/api/idm/resource", hdl.CreateResourceEndpoint).Methods("POST")
	router.HandleFunc("/api/idm/resource/{id}", hdl.UpdateResourceEndpoint).Methods("PUT")

	fmt.Println("Starting server on port 8000...")
	// log.Fatal(http.ListenAndServe(":8000", router))

	corsAllowedOriginsObj := handlers.AllowedOrigins([]string{"*"})
	corsAllowedHeadersObj := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	corsExposedHeadersObj := handlers.ExposedHeaders([]string{"Pagination-Count", "Pagination-Limit", "Pagination-Page"})
	corsAllowedMethodsObj := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(corsAllowedOriginsObj, corsAllowedHeadersObj, corsExposedHeadersObj, corsAllowedMethodsObj)(router)))
}
