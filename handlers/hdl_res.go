package handlers

import (
	"encoding/json"
	"logika/idm/db"
	"logika/idm/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// CreateResourceEndpoint create a resource
func CreateResourceEndpoint(w http.ResponseWriter, r *http.Request) {
	var resource models.Resource
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&resource)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	resourceExist := db.GetResourceByName(resource.Name)
	if resourceExist.Name != "" {
		msg := "Resource already exist in the database"
		http.Error(w, msg, 400)
		return
	}

	resource.ID = primitive.NewObjectID()

	db.CreateResource(resource)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resource)

}

// UpdateResourceEndpoint update a resource
func UpdateResourceEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var resource models.Resource
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&resource)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	db.UpdateResource(oid, resource)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resource)

}

// GetResourceByNameEndpoint get a resource
func GetResourceByNameEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	var resource models.Resource
	_ = json.NewDecoder(r.Body).Decode(&resource)

	resource = db.GetResourceByName(name)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resource)

}

// GetResourceEndpoint get a resource
func GetResourceEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var resource models.Resource
	_ = json.NewDecoder(r.Body).Decode(&resource)

	resource = db.GetResource(oid)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resource)

}

// GetAllResourceEndpoint get a resource
func GetAllResourceEndpoint(w http.ResponseWriter, r *http.Request) {

	var resources []models.Resource
	resources = db.GetAllResource()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resources)

}

// GetPagingResourceEndpoint get a resource
func GetPagingResourceEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slimit := params["limit"]
	soffset := params["offset"]

	// parser limit to int
	limit, err := strconv.ParseInt(slimit, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// parser offset to int
	offset, err := strconv.ParseInt(soffset, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	count := db.CountResource()
	scount := strconv.FormatInt(count, 10)

	page := float64(count) / float64(limit)
	page = math.Ceil(page)
	spage := strconv.FormatFloat(page, 'f', 0, 64)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pagination-Count", scount)
	w.Header().Set("Pagination-Page", spage)
	w.Header().Set("Pagination-Limit", slimit)

	var resources []models.Resource
	resources = db.GetLimitResource(offset, limit)
	json.NewEncoder(w).Encode(resources)

}

// SearchResourceEndpoint search resource
func SearchResourceEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	text := params["text"]
	slimit := params["limit"]
	soffset := params["offset"]

	// parser limit to int
	limit, err := strconv.ParseInt(slimit, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// parser offset to int
	offset, err := strconv.ParseInt(soffset, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var resources []models.Resource
	resources = db.SearchResource(text, offset, limit)

	count := db.SearchResourceCount(text)
	scount := strconv.FormatInt(count, 10)

	page := float64(count) / float64(limit)
	page = math.Ceil(page)
	spage := strconv.FormatFloat(page, 'f', 0, 64)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pagination-Count", scount)
	w.Header().Set("Pagination-Page", spage)
	w.Header().Set("Pagination-Limit", slimit)

	json.NewEncoder(w).Encode(resources)

}

// CoutResourceEndpoint get a resource
func CoutResourceEndpoint(w http.ResponseWriter, r *http.Request) {

	var count int64
	count = db.CountResource()
	res := map[string]int64{"count": count}
	json.NewEncoder(w).Encode(res)

}
