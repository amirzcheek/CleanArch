package main

import (
	"architecture_go/pkg/store/postgres"
	delivery "architecture_go/services/contact/internal/delivery/http"
	repository "architecture_go/services/contact/internal/repository/storage/postgres"
	contactUseCase "architecture_go/services/contact/internal/useCase/contact"
	groupUseCase "architecture_go/services/contact/internal/useCase/group"
	"fmt"
	"log"
	"net/http"
)

func main() {
	dbInfo := &postgres.DBInfo{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "1",
		DBName:   "cleanArch",
	}

	db, err := postgres.ConnectDB(dbInfo)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	fmt.Println("Подключение к БД успешно")

	r := repository.New(db)

	ucContact := contactUseCase.New(r)
	ucGroup := groupUseCase.New(r)

	d := delivery.New(ucContact, ucGroup)

	addr := 4000
	addrStr := fmt.Sprintf(`:%d`, addr)

	log.Printf("Starting server on port: %d", addr)

	if err := http.ListenAndServe(addrStr, d.router); err != nil {
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
