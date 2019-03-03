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

// CreateResource create new resource
func CreateResource(resource models.Resource) (string, error) {
	log.Println("Create a res:", resource.Name)

	res, err := db.Collection("resources").InsertOne(context.TODO(), resource)
	if err != nil {
		log.Print(err)
		return "", err
	}

	log.Println("Inserted res id: ", res.InsertedID)
	return fmt.Sprintf("%v", res.InsertedID), nil
	// fmt.Println("Inserted user: ", user.Username)
}

// UpdateResource update resource
func UpdateResource(id primitive.ObjectID, resource models.Resource) (int64, error) {
	log.Println("Update res:", resource.Name)

	filter := bson.M{"_id": id}
	data := bson.D{
		{"$set", bson.D{
			{"name", resource.Name},
			{"uri", resource.URI},
		}},
		{"$currentDate", bson.D{
			{"lastModified", true},
		}},
	}

	res, err := db.Collection("resources").UpdateOne(context.TODO(), filter, data)
	if err != nil {
		log.Print(err)
		return 0, err
	}

	log.Println("Updated res id: ", res.UpsertedID)
	return res.UpsertedCount, nil
}

// DeleteResource delete resource
func DeleteResource(id primitive.ObjectID) int64 {
	log.Println("Delete resource: ", id)

	filter := bson.M{"_id": id}
	res, err := db.Collection("resources").DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Print(err)
	}

	log.Println("Delete Count : ", res.DeletedCount)
	return res.DeletedCount
}

// GetResource get resource
func GetResource(id primitive.ObjectID) models.Resource {
	log.Println("Get user:", id)

	filter := bson.M{"_id": id}
	var resource models.Resource
	err := db.Collection("resources").FindOne(context.TODO(), filter).Decode(&resource)
	if err != nil {
		log.Print(err)
	}

	log.Println("Get resource: ", resource.Name)
	return resource
}

// GetResourceByName get resource by resource name
func GetResourceByName(name string) models.Resource {
	log.Println("Get resource:", name)

	filter := bson.M{"name": name}
	var resource models.Resource
	err := db.Collection("resources").FindOne(context.TODO(), filter).Decode(&resource)
	if err != nil {
		log.Print(err)
	}

	log.Println("Get resource: ", resource.Name)
	return resource
}

// GetAllResource get all resource
func GetAllResource() []models.Resource {
	log.Println("Get all resource")

	var resources []models.Resource
	cur, err := db.Collection("resources").Find(context.TODO(), bson.D{})
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var resource models.Resource
			err := cur.Decode(&resource)
			if err != nil {
				log.Print(err)
			} else {
				resources = append(resources, resource)
			}

		}
	}

	log.Println("Return all resource: ", resources)
	return resources
}

// GetLimitResource limited resource
func GetLimitResource(offset int64, limit int64) []models.Resource {
	log.Println("Get limit resource")

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	var resources []models.Resource
	cur, err := db.Collection("resources").Find(context.TODO(), bson.D{}, options)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var resource models.Resource
			err := cur.Decode(&resource)
			if err != nil {
				log.Print(err)
			} else {
				resources = append(resources, resource)
			}

		}
	}

	log.Println("Return all resource ")
	return resources
}

// CountResource cout all resource
func CountResource() int64 {
	log.Println("Count all resource")

	cnt, err := db.Collection("resources").Count(context.TODO(), bson.D{})
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	}

	log.Println("Count all resource: ", cnt)
	return cnt
}

// SearchResource search resource
func SearchResource(text string, offset int64, limit int64) []models.Resource {
	log.Println("Search resource")

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	filter := bson.M{"$text": bson.M{"$search": text}}

	var resources []models.Resource
	cur, err := db.Collection("resources").Find(context.TODO(), filter, options)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var resource models.Resource
			err := cur.Decode(&resource)
			if err != nil {
				log.Print(err)
			} else {
				resources = append(resources, resource)
			}
			// log.Print(resource)
		}
	}

	log.Println("Return all resource")
	return resources
}

// SearchResourceCount cout resource
func SearchResourceCount(text string) int64 {
	log.Println("Count search resource")

	filter := bson.M{"$text": bson.M{"$search": text}}
	cnt, err := db.Collection("resources").Count(context.TODO(), filter)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	}

	log.Println("Count search resource: ", cnt)
	return cnt
}
