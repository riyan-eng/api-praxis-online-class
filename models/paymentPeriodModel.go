package models

type PaymentPeriod struct {
	ID         string  `bson:"_id" json:"id"`
	PeriodCode string  `bson:"period_code" json:"period_code"`
	PeriodName string  `bson:"period_name" json:"period_name"`
	Discount   float32 `bson:"discount" json:"discount"`
	IsActive   bool    `bson:"is_active" json:"is_active"`
}
