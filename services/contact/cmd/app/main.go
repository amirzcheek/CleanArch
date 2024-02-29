package main

import (
	"architecture_go/pkg/store/postgres"
	delivery "architecture_go/services/contact/internal/delivery/http"
	repository "architecture_go/services/contact/internal/repository/storage/postgres"
	contactUseCase "architecture_go/services/contact/internal/useCase/contact"
	contactGroupUseCase "architecture_go/services/contact/internal/useCase/contactInGroup"
	groupUseCase "architecture_go/services/contact/internal/useCase/group"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Connect to the database
	dbInfo := &postgres.DBInfo{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "1",
		DBName:   "cleanArch",
	}
	db, err := postgres.ConnectDB(dbInfo)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	fmt.Println("Database connection successful")

	// Initialize repository
	r := repository.New(db)

	// Initialize use cases
	ucContact := contactUseCase.New(r)
	ucGroup := groupUseCase.New(r)
	ucContactGroup := contactGroupUseCase.New(r)

	// Initialize HTTP delivery
	d := delivery.New(ucContact, ucGroup, ucContactGroup)

	// Start HTTP server
	addr := ":4000"
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, d.Router); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
	//firstName, err := name.New("Amirkhan")
	//if err != nil {
	//	fmt.Printf("Error creating first name: %s\n", err)
	//	return
	//}
	//lastName, err := surname.New("Ilyassov")
	//if err != nil {
	//	fmt.Printf("Error creating surname: %s\n", err)
	//	return
	//}
	//phoneNumber, err := phoneNumber.New("87786361157")
	//if err != nil {
	//	fmt.Printf("Error creating phone number: %s\n", err)
	//	return
	//}
	//email, err := email.New("amirzilyassov@gmail.com")
	//if err != nil {
	//	fmt.Printf("Error creating email: %s\n", err)
	//	return
	//}
	//gender, err := gender.NewGender("male")
	//if err != nil {
	//	fmt.Printf("Error creating gender: %s\n", err)
	//	return
	//}
	//
	//amirkhan := contact.NewContact(*firstName, *lastName, *phoneNumber, *email, *gender)
	//fmt.Printf("ID: %s\nName: %s\nSurname: %s\nPhone Number: %s\nEmail: %s\nGender: %s\n", amirkhan.ID, amirkhan.Name, amirkhan.Surname, amirkhan.PhoneNumber, amirkhan.Email, amirkhan.Gender)

}
