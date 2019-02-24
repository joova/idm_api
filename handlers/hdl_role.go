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

// CreateRoleEndpoint create a role
func CreateRoleEndpoint(w http.ResponseWriter, r *http.Request) {
	var role models.Role
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&role)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	roleExist := db.GetRoleByName(role.Name)
	if roleExist.Name != "" {
		msg := "Role already exist in the database"
		http.Error(w, msg, 400)
		return
	}

	role.ID = primitive.NewObjectID()

	db.CreateRole(role)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)

}

// UpdateRoleEndpoint update a role
func UpdateRoleEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var role models.Role
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&role)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	db.UpdateRole(oid, role)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)

}

// GetRoleByNameEndpoint get a role
func GetRoleByNameEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	var role models.Role
	_ = json.NewDecoder(r.Body).Decode(&role)

	role = db.GetRoleByName(name)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)

}

// GetRoleEndpoint get a role
func GetRoleEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var role models.Role
	_ = json.NewDecoder(r.Body).Decode(&role)

	role = db.GetRole(oid)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)

}

// GetAllRoleEndpoint get a role
func GetAllRoleEndpoint(w http.ResponseWriter, r *http.Request) {

	var roles []models.Role
	roles = db.GetAllRole()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)

}

// GetPagingRoleEndpoint get a role
func GetPagingRoleEndpoint(w http.ResponseWriter, r *http.Request) {
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

	count := db.CountRole()
	scount := strconv.FormatInt(count, 10)

	page := float64(count) / float64(limit)
	page = math.Ceil(page)
	spage := strconv.FormatFloat(page, 'f', 0, 64)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pagination-Count", scount)
	w.Header().Set("Pagination-Page", spage)
	w.Header().Set("Pagination-Limit", slimit)

	var roles []models.Role
	roles = db.GetLimitRole(offset, limit)
	json.NewEncoder(w).Encode(roles)

}

// SearchRoleEndpoint search role
func SearchRoleEndpoint(w http.ResponseWriter, r *http.Request) {
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

	var roles []models.Role
	roles = db.SearchRole(text, offset, limit)

	count := db.SearchRoleCount(text)
	scount := strconv.FormatInt(count, 10)

	page := float64(count) / float64(limit)
	page = math.Ceil(page)
	spage := strconv.FormatFloat(page, 'f', 0, 64)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pagination-Count", scount)
	w.Header().Set("Pagination-Page", spage)
	w.Header().Set("Pagination-Limit", slimit)

	json.NewEncoder(w).Encode(roles)

}

// CoutRoleEndpoint get a role
func CoutRoleEndpoint(w http.ResponseWriter, r *http.Request) {

	var count int64
	count = db.CountRole()
	res := map[string]int64{"count": count}
	json.NewEncoder(w).Encode(res)

}
