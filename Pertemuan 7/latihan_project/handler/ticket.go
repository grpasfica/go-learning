package handler

import (
	"fmt"
	"gobus/dto"
	"gobus/service"
)

type TicketHandler struct {
	Service service.TicketService
}

func (h *TicketHandler) CheckTicket(req dto.Request) {
	// Panggil service untuk dapatkan data tiket
	ticket := h.Service.GetTicketPrice(req)

	// Tampilkan output sesuai format di Slide 6
	fmt.Println("=== Harga Tiket ===")
	fmt.Printf("Penumpang : %s\n", ticket.PassengerName)
	fmt.Printf("Tujuan    : %s\n", ticket.Destination)
	// Format %.2f digunakan untuk menampilkan 2 angka desimal (Rp 150000.00)
	fmt.Printf("Harga     : Rp %.2f\n", ticket.Price)
}
