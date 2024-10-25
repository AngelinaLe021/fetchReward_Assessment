package models

type Purchases struct {
	Retailer string  `json:"retailer"` 
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items  []Item `json:"items"`
	Total string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price string `json:"price"`
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
