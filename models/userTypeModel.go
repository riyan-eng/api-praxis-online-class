package models

import (
	"github.com/riyan-eng/api-praxis-online-class/initializers"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserType struct {
	ID           string `json:"id" bson:"_id"`
	UserTypeName string `json:"user_type_name" bson:"user_type_name" validate:"required"`
	UserTypeCode string `json:"user_type_code" bson:"user_type_code" validate:"required"`
	IsActive     bool   `json:"is_active" bson:"is_active"`
}

func UserTypeCollection() *mongo.Collection {
	collection := initializers.Collection("user_types")
	return collection
}
