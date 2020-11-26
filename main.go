package main

import (
	"HWDB/models"
	"HWDB/db"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main()  {
	currency := models.Currency{
		ID:   3,
		Name: "Dollar",
	}

	account := models.Account{
		ID:         3,
		UserId:     3,
		Number:     "125",
		Amount:     1000,
		CurrencyID: 3,
	}
	user := models.User{
		ID:       4,
		Name:     "Mir",
		Surname:  "Mirov",
		Age:      40,
		Gender:   "man",
		Login:    "qwwew",
		Password: "1110",
		Remove:   false,
	}



	DB, err := sql.Open("sqlite3", "bank")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to db")


	err = db.AddNewClient( DB, user)
	if err != nil {
		fmt.Println(err)
	}else {
	fmt.Println("user added")
	}



	err = user.AddNewAccount(DB,account)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("account added")
	}
	err = db.AddNewCurrency(DB, currency)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("currency added")
	}

	/*err = db.RemoveById(DB, 2)
	if err != nil {
		fmt.Println(err)
	}else {
	fmt.Println("user removed")
	}*/


}
