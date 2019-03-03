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

// CreatePrivilegeEndpoint create a privilege
func CreatePrivilegeEndpoint(w http.ResponseWriter, r *http.Request) {
	var privilege models.Privilege
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&privilege)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	privilegeExist := db.GetPrivilegeByName(privilege.Resource.Name)
	if privilegeExist.Resource.Name != "" {
		msg := "Privilege already exist in the database"
		http.Error(w, msg, 400)
		return
	}

	privilege.ID = primitive.NewObjectID()

	db.CreatePrivilege(privilege)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(privilege)

}

// UpdatePrivilegeEndpoint update a privilege
func UpdatePrivilegeEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var privilege models.Privilege
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&privilege)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	db.UpdatePrivilege(oid, privilege)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(privilege)

}

// GetPrivilegeByNameEndpoint get a privilege
func GetPrivilegeByNameEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	var privilege models.Privilege
	_ = json.NewDecoder(r.Body).Decode(&privilege)

	privilege = db.GetPrivilegeByName(name)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(privilege)

}

// DeletePrivilegeEndpoint get a ptype
func DeletePrivilegeEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// _ = json.NewDecoder(r.Body).Decode(&pcat)

	count := db.DeletePrivilege(oid)
	res := map[string]int64{"deleted": count}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

// GetPrivilegeEndpoint get a privilege
func GetPrivilegeEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var privilege models.Privilege
	_ = json.NewDecoder(r.Body).Decode(&privilege)

	privilege = db.GetPrivilege(oid)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(privilege)

}

// GetAllPrivilegeEndpoint get a privilege
func GetAllPrivilegeEndpoint(w http.ResponseWriter, r *http.Request) {

	var privileges []models.Privilege
	privileges = db.GetAllPrivilege()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(privileges)

}

// GetPagingPrivilegeEndpoint get a privilege
func GetPagingPrivilegeEndpoint(w http.ResponseWriter, r *http.Request) {
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

	count := db.CountPrivilege()
	scount := strconv.FormatInt(count, 10)

	page := float64(count) / float64(limit)
	page = math.Ceil(page)
	spage := strconv.FormatFloat(page, 'f', 0, 64)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pagination-Count", scount)
	w.Header().Set("Pagination-Page", spage)
	w.Header().Set("Pagination-Limit", slimit)

	var privileges []models.Privilege
	privileges = db.GetLimitPrivilege(offset, limit)
	json.NewEncoder(w).Encode(privileges)

}

// SearchPrivilegeEndpoint search privilege
func SearchPrivilegeEndpoint(w http.ResponseWriter, r *http.Request) {
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

	var privileges []models.Privilege
	privileges = db.SearchPrivilege(text, offset, limit)

	count := db.SearchPrivilegeCount(text)
	scount := strconv.FormatInt(count, 10)

	page := float64(count) / float64(limit)
	page = math.Ceil(page)
	spage := strconv.FormatFloat(page, 'f', 0, 64)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pagination-Count", scount)
	w.Header().Set("Pagination-Page", spage)
	w.Header().Set("Pagination-Limit", slimit)

	json.NewEncoder(w).Encode(privileges)

}

// CoutPrivilegeEndpoint get a privilege
func CoutPrivilegeEndpoint(w http.ResponseWriter, r *http.Request) {

	var count int64
	count = db.CountPrivilege()
	res := map[string]int64{"count": count}
	json.NewEncoder(w).Encode(res)

}
