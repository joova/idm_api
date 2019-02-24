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

// CreatePrivilege create new privilege
func CreatePrivilege(privilege models.Privilege) (string, error) {
	log.Println("Create a privilege:", privilege.Resource.Name)

	res, err := db.Collection("privileges").InsertOne(context.TODO(), privilege)
	if err != nil {
		log.Print(err)
		return "", err
	}

	log.Println("Inserted res id: ", res.InsertedID)
	return fmt.Sprintf("%v", res.InsertedID), nil
	// fmt.Println("Inserted user: ", user.Username)
}

// UpdatePrivilege update privilege
func UpdatePrivilege(id primitive.ObjectID, privilege models.Privilege) (int64, error) {
	log.Println("Update privilege:", privilege.Resource.Name)

	filter := bson.M{"_id": id}
	data := bson.D{
		{"$set", bson.D{
			{"resource", privilege.Resource},
			{"action", privilege.Action},
		}},
		{"$currentDate", bson.D{
			{"lastModified", true},
		}},
	}

	res, err := db.Collection("privileges").UpdateOne(context.TODO(), filter, data)
	if err != nil {
		log.Print(err)
		return 0, err
	}

	log.Println("Updated res id: ", res.UpsertedID)
	return res.UpsertedCount, nil
}

// GetPrivilege get privilege
func GetPrivilege(id primitive.ObjectID) models.Privilege {
	log.Println("Get privilege:", id)

	filter := bson.M{"_id": id}
	var privilege models.Privilege
	err := db.Collection("privileges").FindOne(context.TODO(), filter).Decode(&privilege)
	if err != nil {
		log.Print(err)
	}

	log.Println("Get privilege: ", privilege.Resource.Name)
	return privilege
}

// GetPrivilegeByName get privilege by privilege name
func GetPrivilegeByName(name string) models.Privilege {
	log.Println("Get privilege:", name)

	filter := bson.M{"name": name}
	var privilege models.Privilege
	err := db.Collection("privileges").FindOne(context.TODO(), filter).Decode(&privilege)
	if err != nil {
		log.Print(err)
	}

	log.Println("Get privilege: ", privilege.Resource.Name)
	return privilege
}

// GetAllPrivilege get all privilege
func GetAllPrivilege() []models.Privilege {
	log.Println("Get all privilege")

	var privileges []models.Privilege
	cur, err := db.Collection("privileges").Find(context.TODO(), bson.D{})
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var privilege models.Privilege
			err := cur.Decode(&privilege)
			if err != nil {
				log.Print(err)
			} else {
				privileges = append(privileges, privilege)
			}

		}
	}

	log.Println("Return all privilege: ", privileges)
	return privileges
}

// GetLimitPrivilege limited privilege
func GetLimitPrivilege(offset int64, limit int64) []models.Privilege {
	log.Println("Get limit privilege")

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	var privileges []models.Privilege
	cur, err := db.Collection("privileges").Find(context.TODO(), bson.D{}, options)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var privilege models.Privilege
			err := cur.Decode(&privilege)
			if err != nil {
				log.Print(err)
			} else {
				privileges = append(privileges, privilege)
			}

		}
	}

	log.Println("Return all privilege ")
	return privileges
}

// CountPrivilege cout all privilege
func CountPrivilege() int64 {
	log.Println("Count all privilege")

	cnt, err := db.Collection("privileges").Count(context.TODO(), bson.D{})
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	}

	log.Println("Count all privilege: ", cnt)
	return cnt
}

// SearchPrivilege search privilege
func SearchPrivilege(text string, offset int64, limit int64) []models.Privilege {
	log.Println("Search privilege")

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	filter := bson.M{"$text": bson.M{"$search": text}}

	var privileges []models.Privilege
	cur, err := db.Collection("privileges").Find(context.TODO(), filter, options)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var privilege models.Privilege
			err := cur.Decode(&privilege)
			if err != nil {
				log.Print(err)
			} else {
				privileges = append(privileges, privilege)
			}
			// log.Print(privilege)
		}
	}

	log.Println("Return all privilege")
	return privileges
}

// SearchPrivilegeCount cout privilege
func SearchPrivilegeCount(text string) int64 {
	log.Println("Count search privilege")

	filter := bson.M{"$text": bson.M{"$search": text}}
	cnt, err := db.Collection("privileges").Count(context.TODO(), filter)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	}

	log.Println("Count search privilege: ", cnt)
	return cnt
}
