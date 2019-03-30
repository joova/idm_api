package db

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo/options"

	"github.com/mongodb/mongo-go-driver/bson"

	"logika/idm/models"
)

// CreateAction create new action
func CreateAction(action models.Action) (string, error) {
	log.Println("Create a action:", action.Name)

	res, err := db.Collection("actions").InsertOne(context.TODO(), action)
	if err != nil {
		log.Print(err)
		return "", err
	}

	log.Println("Inserted res id: ", res.InsertedID)
	return fmt.Sprintf("%v", res.InsertedID), nil
	// fmt.Println("Inserted user: ", user.Username)
}

// UpdateAction update action
func UpdateAction(code string, action models.Action) (int64, error) {
	log.Println("Update action:", action.Name)

	filter := bson.M{"code": code}
	data := bson.D{
		{"$set", bson.D{
			{"name", action.Name},
		}},
		{"$currentDate", bson.D{
			{"lastModified", true},
		}},
	}

	res, err := db.Collection("actions").UpdateOne(context.TODO(), filter, data)
	if err != nil {
		log.Print(err)
		return 0, err
	}

	log.Println("Updated res id: ", res.UpsertedID)
	return res.UpsertedCount, nil
}

// GetAction get action
func GetAction(code string) models.Action {
	log.Println("Get action:", code)

	filter := bson.M{"code": code}
	var action models.Action
	err := db.Collection("actions").FindOne(context.TODO(), filter).Decode(&action)
	if err != nil {
		log.Print(err)
	}

	log.Println("Get action: ", action.Name)
	return action
}

// GetActionByName get action by action name
func GetActionByName(name string) models.Action {
	log.Println("Get action:", name)

	filter := bson.M{"name": name}
	var action models.Action
	err := db.Collection("actions").FindOne(context.TODO(), filter).Decode(&action)
	if err != nil {
		log.Print(err)
	}

	log.Println("Get action: ", action.Name)
	return action
}

// GetAllAction get all action
func GetAllAction() []models.Action {
	log.Println("Get all action")

	var actions []models.Action
	cur, err := db.Collection("actions").Find(context.TODO(), bson.D{})
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var action models.Action
			err := cur.Decode(&action)
			if err != nil {
				log.Print(err)
			} else {
				actions = append(actions, action)
			}

		}
	}

	log.Println("Return all action: ", actions)
	return actions
}

// GetLimitAction limited action
func GetLimitAction(offset int64, limit int64) []models.Action {
	log.Println("Get limit action")

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	var actions []models.Action
	cur, err := db.Collection("actions").Find(context.TODO(), bson.D{}, options)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var action models.Action
			err := cur.Decode(&action)
			if err != nil {
				log.Print(err)
			} else {
				actions = append(actions, action)
			}

		}
	}

	log.Println("Return all action ")
	return actions
}

// CountAction cout all action
func CountAction() int64 {
	log.Println("Count all action")

	cnt, err := db.Collection("actions").Count(context.TODO(), bson.D{})
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	}

	log.Println("Count all action: ", cnt)
	return cnt
}

// SearchAction search action
func SearchAction(text string, offset int64, limit int64) []models.Action {
	log.Println("Search action")

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(offset)

	filter := bson.M{"$text": bson.M{"$search": text}}

	var actions []models.Action
	cur, err := db.Collection("actions").Find(context.TODO(), filter, options)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	} else {
		for cur.Next(context.TODO()) {
			var action models.Action
			err := cur.Decode(&action)
			if err != nil {
				log.Print(err)
			} else {
				actions = append(actions, action)
			}
			// log.Print(action)
		}
	}

	log.Println("Return all action")
	return actions
}

// SearchActionCount cout action
func SearchActionCount(text string) int64 {
	log.Println("Count search action")

	filter := bson.M{"$text": bson.M{"$search": text}}
	cnt, err := db.Collection("actions").Count(context.TODO(), filter)
	// defer cur.Close(context.TODO())
	if err != nil {
		log.Print(err)
	}

	log.Println("Count search action: ", cnt)
	return cnt
}
