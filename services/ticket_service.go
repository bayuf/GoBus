package services

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/bayuf/GoBus/model"
)

type TicketService struct{}

func NewTicketServices() TicketService {
	return TicketService{}
}

// mendapatkan informasi tiket
func (tikcet *TicketService) GetTicket(name, destination string) (model.Tikcet, error) {

	// jika nama kosong
	if name == "" {
		return model.Tikcet{}, errors.New("nama tidak boleh kosong")
	}
	// jika destinasi kosong
	if destination == "" {
		return model.Tikcet{}, errors.New("destinasi tidak boleh kosong")
	}
	// Baca File Json
	dataBytes, err := os.ReadFile("data/destination.json")
	if err != nil {
		return model.Tikcet{}, err
	}

	// decode json => Map
	prices := make(map[string]float64) //membuat map dengan key string dan value float64
	if err := json.Unmarshal(dataBytes, &prices); err != nil {
		return model.Tikcet{}, err
	}

	// mencari harga berdasarkan tujuan di dalam map
	price, ok := prices[destination]
	if !ok {
		return model.Tikcet{}, errors.New("destinasi tidak ditemukan")
	}

	return model.Tikcet{
		Name:        name,
		Destination: destination,
		Price:       price,
	}, nil

}
