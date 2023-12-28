package common

import "go.mongodb.org/mongo-driver/bson/primitive"

type ObjectId struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}
