package handlers

import (
	"encoding/json"
	"logika/idm/crypto"
	"logika/idm/db"
	"logika/idm/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// CreateUserEndpoint create a user
func CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	userExist := db.GetUserByUsername(user.Username)
	if userExist.Username != "" {
		msg := "User already exist in the database"
		http.Error(w, msg, 400)
		return
	}

	user.ID = primitive.NewObjectID()
	hash, _ := crypto.HashPassword(user.Password)
	user.Password = hash

	db.CreateUser(user)

	user.Password = "***"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

// UpdateUserEndpoint update a user
func UpdateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var user models.User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	hash, _ := crypto.HashPassword(user.Password)
	user.Password = hash

	db.UpdateUser(oid, user)
	user.Password = "*******"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

// GetUserByUsernameEndpoint get a user
func GetUserByUsernameEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	user = db.GetUserByUsername(username)
	user.Password = "*******"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

// GetUserEndpoint get a user
func GetUserEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	user = db.GetUser(oid)
	user.Password = "***"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

// GetAllUserEndpoint get a user
func GetAllUserEndpoint(w http.ResponseWriter, r *http.Request) {

	var users []models.User
	users = db.GetAllUser()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}

// GetPagingUserEndpoint get a user
func GetPagingUserEndpoint(w http.ResponseWriter, r *http.Request) {
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

	count := db.CountUser()
	scount := strconv.FormatInt(count, 10)

	page := float64(count) / float64(limit)
	page = math.Ceil(page)
	spage := strconv.FormatFloat(page, 'f', 0, 64)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pagination-Count", scount)
	w.Header().Set("Pagination-Page", spage)
	w.Header().Set("Pagination-Limit", slimit)

	var users []models.User
	users = db.GetLimitUser(offset, limit)
	json.NewEncoder(w).Encode(users)

}

// SearchUserEndpoint search user
func SearchUserEndpoint(w http.ResponseWriter, r *http.Request) {
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

	var users []models.User
	users = db.SearchUser(text, offset, limit)

	count := db.SearchUserCount(text)
	scount := strconv.FormatInt(count, 10)

	page := float64(count) / float64(limit)
	page = math.Ceil(page)
	spage := strconv.FormatFloat(page, 'f', 0, 64)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pagination-Count", scount)
	w.Header().Set("Pagination-Page", spage)
	w.Header().Set("Pagination-Limit", slimit)

	json.NewEncoder(w).Encode(users)

}

// CoutUserEndpoint get a user
func CoutUserEndpoint(w http.ResponseWriter, r *http.Request) {

	var count int64
	count = db.CountUser()
	res := map[string]int64{"count": count}
	json.NewEncoder(w).Encode(res)

}
