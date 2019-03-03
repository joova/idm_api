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

// CreateRole create new role
func CreateRole(role models.Role) (string, error) {
	log.Println("Create a role:", role.Name)

	res, err := db.Collection("roles").InsertOne(context.TODO(), role)
	if err != nil {
		log.Print(err)
		return "", err
	}

	log.Println("Inserted res id: ", res.InsertedID)
	return fmt.Sprintf("%v", res.InsertedID), nil
	// fmt.Println("Inserted user: ", user.Username)
}

// UpdateRole update role
func UpdateRole(id primitive.ObjectID, role models.Role) (int64, error) {
	log.Println("Update role:", role.Name)

	filter := bson.M{"_id": id}
	data := bson.D{
		{"$set", bson.D{
			{"name", role.Name},
		}},
		{"$currentDate", bson.D{
			{"lastModified", true},
		}},
	}

	res, err := db.Collection("roles").UpdateOne(context.TODO(), filter, data)
	if err != nil {
		log.Print(err)
		return 0, err
	}

	log.Println("Updated res id: ", res.UpsertedID)
	return res.UpsertedCount, nil
}

// GetRole get role
func GetRole(id primitive.ObjectID) models.Role {
	log.Println("Get role:", id)

	filter := bson.M{"_id": id}
	var role models.Role
	err := db.Collection("roles").FindOne(context.TODO(), filter).Decode(&role)
	if err != nil {
		log.Print(err)
	}

	log.Println("Get role: ", role.Name)
	return role
}

// DeleteRole delete role
func DeleteRole(id primitive.ObjectID) int64 {
	log.Println("Delete role: ", id)

	filter := bson.M{"_id": id}
	res, err := db.Collection("roles").DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Print(err)
	}

	log.Println("Delete Count : ", res.DeletedCount)
	return res.DeletedCount
}

// GetRoleByName get role by role name
func GetRoleByName(name string) models.Role {
	log.Println("Get role:", name)

	filter := bson.M{"name": name}
	var role models.Role
	err := db.Collection("roles").FindOne(context.TODO(), filter).Decode(&role)
	if err != nil {
		log.Print(err)
	}

	log.Println("Get role: ", role.Name)
	return role
}

// GetAllRole get all role
func GetAllRole() []models.Role {
	log.Println("Get all role")

	var roles []models.Role
	cur, err := db.Collection("roles").Find(context.TODO(), bson.D{})
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var role models.Role
			err := cur.Decode(&role)
			if err != nil {
				log.Print(err)
			} else {
				roles = append(roles, role)
			}

		}
	}

	log.Println("Return all role: ", roles)
	return roles
}

// GetLimitRole limited role
func GetLimitRole(offset int64, limit int64) []models.Role {
	log.Println("Get limit role")

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	var roles []models.Role
	cur, err := db.Collection("roles").Find(context.TODO(), bson.D{}, options)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var role models.Role
			err := cur.Decode(&role)
			if err != nil {
				log.Print(err)
			} else {
				roles = append(roles, role)
			}

		}
	}

	log.Println("Return all role ")
	return roles
}

// CountRole cout all role
func CountRole() int64 {
	log.Println("Count all role")

	cnt, err := db.Collection("roles").Count(context.TODO(), bson.D{})
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	}

	log.Println("Count all role: ", cnt)
	return cnt
}

// SearchRole search role
func SearchRole(text string, offset int64, limit int64) []models.Role {
	log.Println("Search role")

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	filter := bson.M{"$text": bson.M{"$search": text}}

	var roles []models.Role
	cur, err := db.Collection("roles").Find(context.TODO(), filter, options)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var role models.Role
			err := cur.Decode(&role)
			if err != nil {
				log.Print(err)
			} else {
				roles = append(roles, role)
			}
			// log.Print(role)
		}
	}

	log.Println("Return all role")
	return roles
}

// SearchRoleCount cout role
func SearchRoleCount(text string) int64 {
	log.Println("Count search role")

	filter := bson.M{"$text": bson.M{"$search": text}}
	cnt, err := db.Collection("roles").Count(context.TODO(), filter)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	}

	log.Println("Count search role: ", cnt)
	return cnt
}
