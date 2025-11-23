package services

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/bayuf/GoBus/dto"
	"github.com/bayuf/GoBus/model"
)

type TicketService struct {
	Destinations map[string]float64
}

func NewTicketService() *TicketService {
	// Baca File Json
	dataBytes, _ := os.ReadFile("data/destination.json")

	destination := make(map[string]float64)
	json.Unmarshal(dataBytes, &destination)
	return &TicketService{Destinations: destination}
}

// mendapatkan informasi tiket
func (tikcet *TicketService) GetTicket(req dto.Request) (model.Ticket, error) {

	// mencari harga berdasarkan tujuan di dalam map
	price, ok := tikcet.Destinations[req.Destination]
	if !ok {
		return model.Ticket{}, errors.New("destination not found")
	}

	return model.Ticket{
		Name:        req.Name,
		Destination: req.Destination,
		Price:       price,
	}, nil

}

func (tikcet *TicketService) AddDestination(req dto.Request) error {

	// jika data di map ada maka error
	if _, exist := tikcet.Destinations[req.Destination]; exist {
		return errors.New("failed : data already exist")
	}

	tikcet.Destinations[req.Destination] = req.Price
	updated, _ := json.MarshalIndent(tikcet.Destinations, "", " ")
	os.WriteFile("data/destination.json", updated, 0644)

	return nil
}
