package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

// User struct to represent the MongoDB document
type User struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name"`
	Age  int                `json:"age"`
}
