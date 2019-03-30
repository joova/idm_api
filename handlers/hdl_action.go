package handlers

import (
	"encoding/json"
	"logika/idm/db"
	"logika/idm/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateActionEndpoint create a action
func CreateActionEndpoint(w http.ResponseWriter, r *http.Request) {
	var action models.Action
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&action)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	actionExist := db.GetActionByName(action.Name)
	if actionExist.Name != "" {
		msg := "Action already exist in the database"
		http.Error(w, msg, 400)
		return
	}

	db.CreateAction(action)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(action)

}

// UpdateActionEndpoint update a action
func UpdateActionEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["id"]

	var action models.Action
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&action)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	db.UpdateAction(code, action)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(action)

}

// GetActionByNameEndpoint get a action
func GetActionByNameEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	var action models.Action
	_ = json.NewDecoder(r.Body).Decode(&action)

	action = db.GetActionByName(name)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(action)

}

// GetActionEndpoint get a action
func GetActionEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["id"]

	var action models.Action
	_ = json.NewDecoder(r.Body).Decode(&action)

	action = db.GetAction(code)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(action)

}

// GetAllActionEndpoint get a action
func GetAllActionEndpoint(w http.ResponseWriter, r *http.Request) {

	var actions []models.Action
	actions = db.GetAllAction()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actions)

}

// GetPagingActionEndpoint get a action
func GetPagingActionEndpoint(w http.ResponseWriter, r *http.Request) {
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

	count := db.CountAction()
	scount := strconv.FormatInt(count, 10)

	page := float64(count) / float64(limit)
	page = math.Ceil(page)
	spage := strconv.FormatFloat(page, 'f', 0, 64)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pagination-Count", scount)
	w.Header().Set("Pagination-Page", spage)
	w.Header().Set("Pagination-Limit", slimit)

	var actions []models.Action
	actions = db.GetLimitAction(offset, limit)
	json.NewEncoder(w).Encode(actions)

}

// SearchActionEndpoint search action
func SearchActionEndpoint(w http.ResponseWriter, r *http.Request) {
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

	var actions []models.Action
	actions = db.SearchAction(text, offset, limit)

	count := db.SearchActionCount(text)
	scount := strconv.FormatInt(count, 10)

	page := float64(count) / float64(limit)
	page = math.Ceil(page)
	spage := strconv.FormatFloat(page, 'f', 0, 64)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pagination-Count", scount)
	w.Header().Set("Pagination-Page", spage)
	w.Header().Set("Pagination-Limit", slimit)

	json.NewEncoder(w).Encode(actions)

}

// CoutActionEndpoint get a action
func CoutActionEndpoint(w http.ResponseWriter, r *http.Request) {

	var count int64
	count = db.CountAction()
	res := map[string]int64{"count": count}
	json.NewEncoder(w).Encode(res)

}
