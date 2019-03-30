package models

import "github.com/mongodb/mongo-go-driver/bson/primitive"

// Organization organization model
type Organization struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name   string             `bson:"name" json:"name"`
	Parent primitive.ObjectID `bson:"parent,omitempty" json:"parent,omitempty"`
}

// User user model
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string             `bson:"username" json:"username"`
	Password  string             `bson:"password" json:"password"`
	FirstName string             `bson:"first_name" json:"first_name"`
	LastName  string             `bson:"last_name" json:"last_name"`
	Org       primitive.ObjectID `bson:"org,omitempty" json:"org,omitempty"`
}

// Role role model
type Role struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string             `bson:"name" json:"name"`
	Privileges []Privilege        `bson:"previleges" json:"previleges"`
	Org        primitive.ObjectID `bson:"org,omitempty" json:"org,omitempty"`
}

// Privilege previlege model
type Privilege struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Resource Resource           `bson:"resource" json:"resource"`
	Action   Action             `bson:"action" json:"action"`
	Org      primitive.ObjectID `bson:"org,omitempty" json:"org,omitempty"`
}

// Action action model
type Action struct {
	ID   string             `bson:"code,omitempty" json:"code,omitempty"`
	Name string             `bson:"name" json:"name"`
	Org  primitive.ObjectID `bson:"org,omitempty" json:"org,omitempty"`
}

// Resource resource model
type Resource struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string             `bson:"name" json:"name"`
	URI  string             `bson:"uri" json:"uri"`
	Org  primitive.ObjectID `bson:"org,omitempty" json:"org,omitempty"`
}
