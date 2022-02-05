package verification

import "github.com/KonstantinGasser/weeat/core/dao"

type Type int

const (
	Unset Type = iota
	LowAck
	MediumAck
	HighAck
)

type FoodEvent struct {
	ID               int
	Status           int
	Acknowledgements int
	Convidence       float64
}

func ToDAO(evt FoodEvent) dao.VerifyItem {
	return dao.VerifyItem{
		ID:               evt.ID,
		Status:           evt.Status,
		Acknowledgements: evt.Acknowledgements,
		Convidence:       evt.Convidence,
	}
}
