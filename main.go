package main

import (
	"fmt"
	"time"
	"sync"
)

type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

const conferenceName string=  "Go Conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = make([]UserData,0)

func main(){

	greet()

	for {

		firstName,lastName,email,userTickets := getUserInput()
		isValidName,isValidEmail,isValidTicketNumber :=isValidInput(firstName, lastName, userTickets,email,remainingTickets)

		if isValidTicketNumber && isValidEmail && isValidName {
			
			bookTicket(userTickets,firstName,lastName,email)

			wg.Add(1)
			go sendTicket(userTickets,firstName,lastName,email)

			getFirstNames()

			noTicketsRemaning := remainingTickets == 0
			if noTicketsRemaning {
				fmt.Println("Our conference is booked out.Come back next year.")
				break
			} 
		}else{
			if !isValidName{
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail{
				fmt.Println("Entered email address doesn't contain @ symbol")
			}
			if !isValidTicketNumber{
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
		wg.Wait()
	}
}

func getFirstNames(){
	firstNames := []string{}
	for _,booking := range bookings {
		firstNames = append(firstNames,booking.firstName)
	}
	fmt.Printf("The first name of bookings are %v\n",firstNames)
}

func bookTicket(userTickets uint,firstName string,lastName string,email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings,userData)
	fmt.Printf("List of bookings is %v\n",bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a conformation email at %v\n",firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets left for %v\n",remainingTickets,conferenceName)
}

func greet(){
	fmt.Printf("Welcome to %v booking application.\n",conferenceName)
	fmt.Println("We have total of",conferenceTickets,"tickets and",remainingTickets,"are still avaliable.")
	fmt.Println("Get your tickets here to attend")
}

func getUserInput()(string, string, string,uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your First name :")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last name :")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address :")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets :")
	fmt.Scan(&userTickets)

	return firstName,lastName,email,userTickets;
}

func sendTicket(userTickets uint,firstName string,lastName string,email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v ticket for %v %v",userTickets,firstName,lastName)
	fmt.Println("###########")
	fmt.Printf("Sending ticket: \n %v to email address %v\n",ticket,email)
	fmt.Println("###########")
	wg.Done()
}