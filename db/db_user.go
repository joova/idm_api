package db

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo/options"

	"github.com/mongodb/mongo-go-driver/bson"

	"logika/idm/models"
)

// CreateUser create new user
func CreateUser(user models.User) (string, error) {
	log.Println("Create a user:", user.Username)

	res, err := db.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		log.Print(err)
		return "", err
	}

	log.Println("Inserted user id: ", res.InsertedID)
	return fmt.Sprintf("%v", res.InsertedID), nil
	// fmt.Println("Inserted user: ", user.Username)
}

// UpdateUser update user
func UpdateUser(id primitive.ObjectID, user models.User) (int64, error) {
	log.Println("Update user:", user.Username)

	filter := bson.M{"_id": id}
	data := bson.D{
		{"$set", bson.D{
			{"username", user.Username},
			{"password", user.Password},
			{"first_name", user.FirstName},
			{"last_name", user.LastName},
		}},
		{"$currentDate", bson.D{
			{"lastModified", true},
		}},
	}

	res, err := db.Collection("users").UpdateOne(context.TODO(), filter, data)
	if err != nil {
		log.Print(err)
		return 0, err
	}

	log.Println("Updated user id: ", res.UpsertedID)
	return res.UpsertedCount, nil
}

// GetUser get user
func GetUser(id primitive.ObjectID) models.User {
	log.Println("Get user:", id)

	filter := bson.M{"_id": id}
	var user models.User
	err := db.Collection("users").FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Print(err)
	}

	log.Println("Get user: ", user.Username)
	return user
}

// GetUserByUsername get user
func GetUserByUsername(username string) models.User {
	log.Println("Get user:", username)

	filter := bson.M{"username": username}
	var user models.User
	err := db.Collection("users").FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Print(err)
	}

	log.Println("Get user: ", user.Username)
	return user
}

// GetAllUser get user
func GetAllUser() []models.User {
	log.Println("Get all user")

	var users []models.User
	cur, err := db.Collection("users").Find(context.TODO(), bson.D{})
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var user models.User
			err := cur.Decode(&user)
			if err != nil {
				log.Print(err)
			} else {
				user.Password = "***"
				users = append(users, user)
			}

		}
	}

	log.Println("Return all user: ", users)
	return users
}

// GetLimitUser get user
func GetLimitUser(offset int64, limit int64) []models.User {
	log.Println("Get limit user")

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	var users []models.User
	cur, err := db.Collection("users").Find(context.TODO(), bson.D{}, options)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var user models.User
			err := cur.Decode(&user)
			if err != nil {
				log.Print(err)
			} else {
				user.Password = "***"
				users = append(users, user)
			}

		}
	}

	log.Println("Return all user: ", users)
	return users
}

// CountUser cout all user
func CountUser() int64 {
	log.Println("Count all user")

	cnt, err := db.Collection("users").Count(context.TODO(), bson.D{})
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	}

	log.Println("Count all user: ", cnt)
	return cnt
}

// SearchUser get user
func SearchUser(text string, offset int64, limit int64) []models.User {
	log.Println("Search user")

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	filter := bson.M{"$text": bson.M{"$search": text}}

	var users []models.User
	cur, err := db.Collection("users").Find(context.TODO(), filter, options)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var user models.User
			err := cur.Decode(&user)
			if err != nil {
				log.Print(err)
			} else {
				user.Password = "***"
				users = append(users, user)
			}
			// log.Print(user)
		}
	}

	log.Println("Return all user: ", users)
	return users
}

// SearchUserCount cout all user
func SearchUserCount(text string) int64 {
	log.Println("Count all user")

	filter := bson.M{"$text": bson.M{"$search": text}}
	cnt, err := db.Collection("users").Count(context.TODO(), filter)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	}

	log.Println("Count all user: ", cnt)
	return cnt
}
