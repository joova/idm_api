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

// CreateOrg create new org
func CreateOrg(org models.Organization) (string, error) {
	log.Println("Create a org:", org.Name)

	res, err := db.Collection("orgs").InsertOne(context.TODO(), org)
	if err != nil {
		log.Print(err)
		return "", err
	}

	log.Println("Inserted org id: ", res.InsertedID)
	return fmt.Sprintf("%v", res.InsertedID), nil
	// fmt.Println("Inserted user: ", user.Username)
}

// UpdateOrg update org
func UpdateOrg(id primitive.ObjectID, org models.Organization) (int64, error) {
	log.Println("Update org:", org.Name)

	filter := bson.M{"_id": id}
	data := bson.D{
		{"$set", bson.D{
			{"name", org.Name},
			{"parent", org.Parent},
		}},
		{"$currentDate", bson.D{
			{"lastModified", true},
		}},
	}

	res, err := db.Collection("orgs").UpdateOne(context.TODO(), filter, data)
	if err != nil {
		log.Print(err)
		return 0, err
	}

	log.Println("Updated org id: ", res.UpsertedID)
	return res.UpsertedCount, nil
}

// GetOrg get org
func GetOrg(id primitive.ObjectID) models.Organization {
	log.Println("Get user:", id)

	filter := bson.M{"_id": id}
	var org models.Organization
	err := db.Collection("orgs").FindOne(context.TODO(), filter).Decode(&org)
	if err != nil {
		log.Print(err)
	}

	log.Println("Get org: ", org.Name)
	return org
}

// GetOrgByName get org by org name
func GetOrgByName(name string) models.Organization {
	log.Println("Get org:", name)

	filter := bson.M{"name": name}
	var org models.Organization
	err := db.Collection("orgs").FindOne(context.TODO(), filter).Decode(&org)
	if err != nil {
		log.Print(err)
	}

	log.Println("Get org: ", org.Name)
	return org
}

// GetAllOrg get all org
func GetAllOrg() []models.Organization {
	log.Println("Get all org")

	var orgs []models.Organization
	cur, err := db.Collection("orgs").Find(context.TODO(), bson.D{})
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var org models.Organization
			err := cur.Decode(&org)
			if err != nil {
				log.Print(err)
			} else {
				orgs = append(orgs, org)
			}

		}
	}

	log.Println("Return all org: ", orgs)
	return orgs
}

// GetLimitOrg limited org
func GetLimitOrg(offset int64, limit int64) []models.Organization {
	log.Println("Get limit org")

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	var orgs []models.Organization
	cur, err := db.Collection("orgs").Find(context.TODO(), bson.D{}, options)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var org models.Organization
			err := cur.Decode(&org)
			if err != nil {
				log.Print(err)
			} else {
				orgs = append(orgs, org)
			}

		}
	}

	log.Println("Return all org ")
	return orgs
}

// CountOrg cout all org
func CountOrg() int64 {
	log.Println("Count all org")

	cnt, err := db.Collection("orgs").Count(context.TODO(), bson.D{})
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	}

	log.Println("Count all org: ", cnt)
	return cnt
}

// SearchOrg search org
func SearchOrg(text string, offset int64, limit int64) []models.Organization {
	log.Println("Search org")

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	filter := bson.M{"$text": bson.M{"$search": text}}

	var orgs []models.Organization
	cur, err := db.Collection("orgs").Find(context.TODO(), filter, options)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var org models.Organization
			err := cur.Decode(&org)
			if err != nil {
				log.Print(err)
			} else {
				orgs = append(orgs, org)
			}
			// log.Print(org)
		}
	}

	log.Println("Return all org")
	return orgs
}

// SearchOrgCount cout org
func SearchOrgCount(text string) int64 {
	log.Println("Count search org")

	filter := bson.M{"$text": bson.M{"$search": text}}
	cnt, err := db.Collection("orgs").Count(context.TODO(), filter)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	}

	log.Println("Count search org: ", cnt)
	return cnt
}
