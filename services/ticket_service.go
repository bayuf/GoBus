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

func NewTicketService() *TicketService {
	return &TicketService{}
}

// mendapatkan informasi tiket
func (ticket *TicketService) GetTicket(req dto.Request) (model.Ticket, error) {
	// buka file json
	file, err := os.Open("data/destination.json")
	if err != nil {
		return model.Ticket{}, err
	}

	// file ditutup setelah semua proses di func ini selesai dijalankan
	defer file.Close()

	// baca dengan streaming
	decoder := json.NewDecoder(file)

	destinations := make(map[string]float64) //map kosong
	decoder.Decode(&destinations)            // decode json menjadi map

	// jika nama kosong
	if strings.TrimSpace(req.Name) == "" {
		return model.Ticket{}, errors.New("error: name is empty")
	}
	// jika destinasi kosong
	if strings.TrimSpace(req.Destination) == "" {
		return model.Ticket{}, errors.New("error: destination is empty")
	}

	// mencari harga berdasarkan tujuan di dalam map
	price, ok := destinations[req.Destination]
	if !ok {
		return model.Ticket{}, errors.New("error: destination not found")
	}

	return model.Ticket{
		Name:        req.Name,
		Destination: req.Destination,
		Price:       price,
	}, nil

}

// func (tikcet *TicketService) AddDestination(req dto.Request) error {
// 	// jika destinasi kosong
// 	if strings.TrimSpace(req.Destination) == "" {
// 		return errors.New("error: destination is empty")
// 	}

// 	// jika biaya 0 atau kurang
// 	if req.Price <= 0 {
// 		return errors.New("error: invalid price input")
// 	}

// 	// validasi jika destinasi sudah ada
// 	if _, exist := tikcet.Destinations[req.Destination]; exist {
// 		return errors.New("error: data already exist")
// 	}

// 	tikcet.Destinations[req.Destination] = req.Price

// 	// membuat file baru untuk menampung data sementara
// 	tempFile := "data/destination_temp.json"

// 	file, err := os.Create(tempFile)
// 	if err != nil {
// 		return err
// 	}

// 	// encode json ke temp file
// 	encoder := json.NewEncoder(file)
// 	encoder.SetIndent("", "  ")

// 	if err := encoder.Encode(tikcet.Destinations); err != nil {
// 		return err
// 	}

// 	defer file.Close() // tutup file sebelum replace
// 	if err := os.Rename(tempFile, "data/destination.json"); err != nil {
// 		return err
// 	}
// 	return nil
// }
