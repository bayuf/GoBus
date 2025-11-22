package handler

import (
	"errors"
	"strings"

	"github.com/bayuf/GoBus/dto"
	"github.com/bayuf/GoBus/services"
)

type Handler struct {
	TicketService services.TicketService
}

func NewHandler(ts *services.TicketService) Handler {
	return Handler{TicketService: *ts}
}

func (handler *Handler) Process(req dto.Request) (dto.Response, error) {

	// jika nama kosong
	if strings.TrimSpace(req.Name) == "" {
		return dto.Response{}, errors.New("name is empty")
	}
	// jika destinasi kosong
	if strings.TrimSpace(req.Destination) == "" {
		return dto.Response{}, errors.New("destination is empty")
	}

	// Request dikirim ke services untuk mendapatkan informasi tiket
	ticket, err := handler.TicketService.GetTicket(req)
	if err != nil {
		return dto.Response{}, err
	}

	return dto.Response{
		Name:        ticket.Name,
		Destination: ticket.Destination,
		Price:       ticket.Price,
	}, nil
}
