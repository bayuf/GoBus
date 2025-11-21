package main

import (
	"fmt"

	"github.com/bayuf/GoBus/dto"
	"github.com/bayuf/GoBus/handler"
)

func main() {
	// mengirim request ke dto
	request := dto.NewRequest("Bayu Firmansyah", "Jakarta")

	response, err := handler.Process(request)
	if err != nil {
		fmt.Println("Data Error :", err.Error())
	} else {
		fmt.Println("======= Harga Tiket =======")
		fmt.Println("Penumpang :", response.Name)
		fmt.Println("Tujuan\t  :", response.Destination)
		fmt.Printf("Harga\t  : Rp. %.2f\n", response.Price)
		fmt.Println("===========================")
	}

}
