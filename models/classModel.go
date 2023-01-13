package models

type Class struct {
	ID              string `json:"id" bson:"_id"`
	ClassName       string `json:"class_name" bson:"class_name"`
	ClassCode       string `json:"class_code" bson:"class_code"`
	ClassMonthPrice int    `json:"class_month_price" bson:"class_month_price"`
	// IsActive        bool   `json:"is_active" bson:"is_active"`
}
