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

// CreateOrgEndpoint create a org
func CreateOrgEndpoint(w http.ResponseWriter, r *http.Request) {
	var org models.Organization
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&org)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	orgExist := db.GetOrgByName(org.Name)
	if orgExist.Name != "" {
		msg := "Org already exist in the database"
		http.Error(w, msg, 400)
		return
	}

	org.ID = primitive.NewObjectID()

	db.CreateOrg(org)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(org)

}

// UpdateOrgEndpoint update a org
func UpdateOrgEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var org models.Organization
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&org)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	db.UpdateOrg(oid, org)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(org)

}

// GetOrgByOrgnameEndpoint get a org
func GetOrgByOrgnameEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	var org models.Organization
	_ = json.NewDecoder(r.Body).Decode(&org)

	org = db.GetOrgByName(name)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(org)

}

// GetOrgEndpoint get a org
func GetOrgEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var org models.Organization
	_ = json.NewDecoder(r.Body).Decode(&org)

	org = db.GetOrg(oid)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(org)

}

// GetAllOrgEndpoint get a org
func GetAllOrgEndpoint(w http.ResponseWriter, r *http.Request) {

	var orgs []models.Organization
	orgs = db.GetAllOrg()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orgs)

}

// GetPagingOrgEndpoint get a org
func GetPagingOrgEndpoint(w http.ResponseWriter, r *http.Request) {
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

	count := db.CountOrg()
	scount := strconv.FormatInt(count, 10)

	page := float64(count) / float64(limit)
	page = math.Ceil(page)
	spage := strconv.FormatFloat(page, 'f', 0, 64)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pagination-Count", scount)
	w.Header().Set("Pagination-Page", spage)
	w.Header().Set("Pagination-Limit", slimit)

	var orgs []models.Organization
	orgs = db.GetLimitOrg(offset, limit)
	json.NewEncoder(w).Encode(orgs)

}

// SearchOrgEndpoint search org
func SearchOrgEndpoint(w http.ResponseWriter, r *http.Request) {
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

	var orgs []models.Organization
	orgs = db.SearchOrg(text, offset, limit)

	count := db.SearchOrgCount(text)
	scount := strconv.FormatInt(count, 10)

	page := float64(count) / float64(limit)
	page = math.Ceil(page)
	spage := strconv.FormatFloat(page, 'f', 0, 64)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pagination-Count", scount)
	w.Header().Set("Pagination-Page", spage)
	w.Header().Set("Pagination-Limit", slimit)

	json.NewEncoder(w).Encode(orgs)

}

// CoutOrgEndpoint get a org
func CoutOrgEndpoint(w http.ResponseWriter, r *http.Request) {

	var count int64
	count = db.CountOrg()
	res := map[string]int64{"count": count}
	json.NewEncoder(w).Encode(res)

}
