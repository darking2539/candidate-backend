package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskModel struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title        string             `json:"title" bson:"title"`
	Description  string             `json:"description" bson:"description"`
	Status       string             `json:"status" bson:"status"`
	KeepStatus   bool               `json:"-" bson:"keepStatus"`
	CreatedEmail string             `json:"createdEmail" bson:"createdEmail"`
	CreatedName  string             `json:"createdName" bson:"createdName"`
	CreatedDated time.Time          `json:"createdDated" bson:"createdDated"`
	Comment      []Comment          `json:"comments" bson:"comments"`
}

type Comment struct {
	CreatedEmail string             `json:"createdEmail" bson:"createdEmail"`
	CreatedName  string             `json:"createdName" bson:"createdName"`
	CommentDesp  string             `json:"commentDesp" bson:"commentDesp"`
	CreatedDated time.Time          `json:"createdDated" bson:"createdDated"`
}