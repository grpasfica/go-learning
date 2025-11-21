package main

import (
	"gobus/dto"
	"gobus/handler"
	"gobus/service"
)

func main() {
	// Setup Dependency
	ticketService := service.TicketService{}
	ticketHandler := handler.TicketHandler{Service: ticketService}

	// 2Request sesuai contoh di Slide 5
	req := dto.NewRequest("graciarp_", "Jakarta")

	// Eksekusi Handler
	ticketHandler.CheckTicket(req)
}
