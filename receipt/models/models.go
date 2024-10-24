package models

import (
	"gorm.io/gorm"
)

type Purchases struct {
	gorm.Model
	// ID uint `gorm:"primaryKey`
	Retailer string  `json:"retailer" gorm:"varchar(225);not null;default:null` 
	PurchaseDate string `json:"purchaseDate" gorm:"type:date"`
	PurchaseTime string `json:"purchaseTime" gorm:"type:time`
	Items  []Item `json:"items" gorm:"foreignKey:PurchaseID`
	Total string `json:"total" gorm:"type:varchar(225);not null;default:0.0`
}

type Item struct {
	gorm.Model
	// ID uint `gorm:"primaryKey"`
	// PurchaseID uint `json:"purchaseId" gorm:not null`
	ShortDescription string `json:"shortDescription" gorm:"type:varchar(225);not null`
	Price string `json:"price" gorm:"type:varchar(225);not null`
}

type Response struct {
	ID string `json:"id"`
	Points int `json:"points"`
}

type Storage struct {
	Store map[string]Response
}

func NewRecord() *Storage {
	return &Storage{
		Store: make(map[string]Response),
	}
}

func (s *Storage) Add(newRecord Response) {
	s.Store[newRecord.ID] = newRecord
}

func (s *Storage) Fetch(id string) (Response, bool) {
	record, exist := s.Store[id]
	return record, exist
}
