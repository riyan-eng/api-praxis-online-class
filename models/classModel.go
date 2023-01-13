package models

type Class struct {
	ID              string `json:"id" bson:"_id"`
	ClassName       string `json:"class_name" bson:"class_name" validate:"required"`
	ClassCode       string `json:"class_code" bson:"class_code" validate:"required"`
	ClassMonthPrice int    `json:"class_month_price" bson:"class_month_price" validate:"required"`
	// IsActive        bool   `json:"is_active" bson:"is_active"`
}
