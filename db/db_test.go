package db

import (
	"logika/idm/crypto"
	"logika/idm/models"
	"testing"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

func TestCreateUser(t *testing.T) {
	var user models.User

	user.ID = primitive.NewObjectID()
	user.Username = "user1@email.com"
	hash, _ := crypto.HashPassword("password1")
	user.Password = hash

	_, err := CreateUser(user)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateUser(t *testing.T) {
	id := "5c5a6b696f3359f247f9d47d"
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Error(err.Error())
	}

	var user models.User
	user.ID = oid
	user.Username = "test2@email.com"

	hash, _ := crypto.HashPassword("passowrd2")
	user.Password = hash

	_, err = UpdateUser(oid, user)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestSearch(t *testing.T) {

	limit := 10
	offset := 0

	count := SearchUserCount("hasyim")
	// scount := strconv.FormatInt(count, 10)

	// page := float64(count) / float64(limit)
	// page = math.Ceil(page)
	// spage := strconv.FormatFloat(page, 'f', 0, 64)

	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Pagination-Count", scount)
	// w.Header().Set("Pagination-Page", spage)
	// w.Header().Set("Pagination-Limit", slimit)

	var users []models.User
	users = SearchUser("hasyim", int64(offset), int64(limit))

	if len(users) != int(count) {
		t.Error("Len user != 2")
	}

}
