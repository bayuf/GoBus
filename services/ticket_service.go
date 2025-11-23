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
	// buka file json
	file, err := os.Open("data/destination.json")
	if err != nil {
		// kembali map kosong jika file tidak ada
		return &TicketService{Destinations: make(map[string]float64)}
	}

	// baca dengan streaming
	decoder := json.NewDecoder(file)

	destinations := make(map[string]float64) //map kosong
	decoder.Decode(&destinations)            // decode json menjadi map

	// file ditutup setelah semua proses di func ini selesai dijalankan
	file.Close()

	return &TicketService{Destinations: destinations}

	// Baca File Json
	// dataBytes, _ := os.ReadFile("data/destination.json")

	// destination := make(map[string]float64)
	// json.Unmarshal(dataBytes, &destination)
	// return &TicketService{Destinations: destination}
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

	// validasi jika destinasi sudah ada
	if _, exist := tikcet.Destinations[req.Destination]; exist {
		return errors.New("failed: data already exist")
	}

	tikcet.Destinations[req.Destination] = req.Price

	// membuat file baru untuk menampung data sementara
	tempFile := "data/destination_temp.json"

	file, err := os.Create(tempFile)
	if err != nil {
		return err
	}

	// encode json ke temp file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(tikcet.Destinations); err != nil {
		return err
	}

	file.Close() // tutup file sebelum replace
	if err := os.Rename(tempFile, "data/destination.json"); err != nil {
		return err
	}

	// jika data di map ada maka error
	// if _, exist := tikcet.Destinations[req.Destination]; exist {
	// 	return errors.New("failed : data already exist")
	// }

	// tikcet.Destinations[req.Destination] = req.Price
	// updated, _ := json.MarshalIndent(tikcet.Destinations, "", "  ")
	// os.WriteFile("data/destination.json", updated, 0644)

	return nil
}
