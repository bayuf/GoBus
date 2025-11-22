package services

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/bayuf/GoBus/dto"
	"github.com/bayuf/GoBus/model"
)

type TicketService struct {
	Price map[string]float64
}

func NewTicketService() *TicketService {
	// Baca File Json
	dataBytes, _ := os.ReadFile("data/destination.json")

	price := make(map[string]float64)
	json.Unmarshal(dataBytes, &price)
	return &TicketService{Price: price}
}

// mendapatkan informasi tiket
func (tikcet *TicketService) GetTicket(req dto.Request) (model.Ticket, error) {

	// mencari harga berdasarkan tujuan di dalam map
	price, ok := tikcet.Price[req.Destination]
	if !ok {
		return model.Ticket{}, errors.New("destination not found")
	}

	return model.Ticket{
		Name:        req.Name,
		Destination: req.Destination,
		Price:       price,
	}, nil

}

// func (tikcet *TicketService) AddDestination(req dto.Request) error {
// 	// jika nama kosong
// 	if strings.TrimSpace(req.Destination) == "" {
// 		return errors.New("destinasi tidak boleh kosong")
// 	}
// 	// jika destinasi kosong
// 	if req.Price <= 0 {
// 		return errors.New("harga tidak valid")
// 	}

// 	return nil
// }
