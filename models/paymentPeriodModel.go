package models

import (
	"github.com/riyan-eng/api-praxis-online-class/initializers"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentPeriod struct {
	ID         string  `bson:"_id" json:"id"`
	PeriodCode string  `bson:"period_code" json:"period_code"`
	PeriodName string  `bson:"period_name" json:"period_name"`
	Discount   float32 `bson:"discount" json:"discount"`
	IsActive   bool    `bson:"is_active" json:"is_active"`
}

func PaymentPeriodCollection() *mongo.Collection {
	collection := initializers.Collection("payment_period")
	return collection
}
