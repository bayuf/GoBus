package handler

import (
	"github.com/bayuf/GoBus/dto"
	"github.com/bayuf/GoBus/services"
)

func Process(req dto.Request) (dto.Response, error) {
	ticket, err := services.GetTicket(req.Name, req.Destination)
	if err != nil {
		return dto.Response{}, err
	}

	return dto.Response{
		Name:        ticket.Name,
		Destination: ticket.Destination,
		Price:       ticket.Price,
	}, nil
}
