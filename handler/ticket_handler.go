package handler

import (
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

func (handler *Handler) AddDest(req dto.Request) error {
	// mengirim request ke service untuk menambah destinasi
	if err := handler.TicketService.AddDestination(req); err != nil {
		return err
	}
	return nil
}
