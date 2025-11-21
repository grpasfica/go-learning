package service

import (
	"gobus/data"
	"gobus/dto"
	"gobus/model"
)

type TicketService struct{}

func (s *TicketService) GetTicketPrice(req dto.Request) model.Ticket {
	// Ambil harga dari package data
	price, exists := data.Destination[req.Destination]

	finalPrice := 0.0
	if exists {
		finalPrice = float64(price)
	}

	return model.Ticket{
		PassengerName: req.Name,
		Destination:   req.Destination,
		Price:         finalPrice,
	}
}
