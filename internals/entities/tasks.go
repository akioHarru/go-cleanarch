package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Created_At  time.Time          `bson:"created_at" json:"created_at"`
	Due_Date    time.Time          `bson:"due_date" json:"due_date"`
	Completed   bool               `bson:"completed" json:"completed"`
}
