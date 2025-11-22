package services

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/bayuf/GoBus/dto"
	"github.com/bayuf/GoBus/model"
)

type TicketService struct{}

func NewTicketService() TicketService {
	return TicketService{}
}

// mendapatkan informasi tiket
func (tikcet *TicketService) GetTicket(req dto.Request) (model.Ticket, error) {

	// jika nama kosong
	if strings.TrimSpace(req.Name) == "" {
		return model.Ticket{}, errors.New("name is empty")
	}
	// jika destinasi kosong
	if strings.TrimSpace(req.Destination) == "" {
		return model.Ticket{}, errors.New("destination is empty")
	}
	// Baca File Json
	dataBytes, err := os.ReadFile("data/destination.json")
	if err != nil {
		return model.Ticket{}, err
	}

	// decode json => Map
	prices := make(map[string]float64) //membuat map dengan key string dan value float64
	if err := json.Unmarshal(dataBytes, &prices); err != nil {
		return model.Ticket{}, err
	}

	// mencari harga berdasarkan tujuan di dalam map
	price, ok := prices[req.Destination]
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
