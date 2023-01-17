package models

import (
	"github.com/riyan-eng/api-praxis-online-class/initializers"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID              string `bson:"_id" json:"id"`
	UserTypeId      string `bson:"user_type" json:"user_type"`
	ClassId         string `bson:"class" json:"class" validate:"required"`
	Name            string `bson:"name" json:"name" validate:"required"`
	Birthday        string `bson:"birthday" json:"birthday" validate:"required"`
	Phone           string `bson:"phone" json:"phone" validate:"required"`
	Email           string `bson:"email" json:"email" validate:"required"`
	Address         string `bson:"address" json:"address" validate:"required"`
	Education       string `bson:"education" json:"education" validate:"required"`
	Reference       string `bson:"reference" json:"reference"`
	PaymentPeriodId string `bson:"payment_period" json:"payment_period" validate:"required"`
	UniqueCode      uint16 `bson:"unique_code" json:"unique_code"`
	Batch           string `bson:"batch" json:"batch" validate:"required"`
	IsActive        bool   `bson:"is_active" json:"is_active"`
}

func UserCollection() *mongo.Collection {
	collection := initializers.Collection("users")
	return collection
}
