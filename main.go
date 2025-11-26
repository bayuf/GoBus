package main

import (
	"fmt"

	"github.com/bayuf/GoBus/dto"
	"github.com/bayuf/GoBus/handler"
	"github.com/bayuf/GoBus/services"
)

func main() {
	// Init
	ticketService := services.NewTicketService()
	handler := handler.NewHandler(ticketService)

	// Menambah destinasi baru
	// addDestReq := dto.NewDest("Mojokerto", 10000)

	// if err := handler.AddDest(addDestReq); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("succes")
	// }

	// mengirim request ke dto
	request := dto.NewRequest("Bayu Firmansyah", "Surabaya")

	response, err := handler.Process(request)
	if err != nil {
		fmt.Println("Data Error :", err)
	} else {
		fmt.Println("======= Harga Tiket =======")
		fmt.Println("Penumpang :", response.Name)
		fmt.Println("Tujuan\t  :", response.Destination)
		fmt.Printf("Harga\t  : Rp. %.2f\n", response.Price)
		fmt.Println("===========================")
	}

}
