package main

import (
	"strings"
)

func isValidInput(firstName string,lastName string,userTickets uint,email string,remainingTickets uint)(bool,bool,bool){
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isVaidEmail := strings.Contains(email,"@")
	isValidTicketNumber := userTickets >0 && userTickets < remainingTickets
	return isValidName,isVaidEmail,isValidTicketNumber;
}

