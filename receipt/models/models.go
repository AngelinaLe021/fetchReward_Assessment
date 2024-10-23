package models

import "gorm.io/gorm"

// type Purchases struct {
// 	gorm.Model

// 	ID uint `gorm:"primaryKey"`
// 	Retailer string  `json:"retailer" gorm:"text;not null;default:null` 
// 	PurchaseDate time.Time `json:"purchaseDate" gorm:"type:date"`
// 	PurchaseTime time.Time `json:"purchaseTime" gorm:"type:time`
// 	Items  []Item `json:"items" gorm:"foreignKey:PurchaseID`
// 	Total float64 `json:"total" gorm:"type:numeric(10,2);not null;default:0.0`
// }

// type Item struct {
// 	gorm.Model
// 	ID uint `gorm:"primaryKey"`
// 	PurchaseID uint `json:"purchaseId" gorm:`
// }

type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"text;not null;default:null`
	Answer string `json:"answer" gorm:"text;not null;default:null`
}