package services

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/bayuf/GoBus/model"
)

// mendapatkan informasi tiket
func GetTicket(name, destination string) (model.Ticket, error) {

	// jika nama kosong
	if name == "" {
		return model.Ticket{}, errors.New("nama tidak boleh kosong")
	}
	// jika destinasi kosong
	if destination == "" {
		return model.Ticket{}, errors.New("destinasi tidak boleh kosong")
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
	price, ok := prices[destination]
	if !ok {
		return model.Ticket{}, errors.New("destinasi tidak ditemukan")
	}

	return model.Ticket{
		Name:        name,
		Destination: destination,
		Price:       price,
	}, nil

}
