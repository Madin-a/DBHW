package db

import (
	"HWDB/models"
	"database/sql"
	"fmt"
)


func AddNewClient(db *sql.DB, client models.User)(err error) {
	_, err = db.Exec("INSERT into user(Name, Surname, Age, Gender, Login, Password, Remove) VALUES (($1),($2),($3),($4),($5), ($6), ($7))", client.Name, client.Surname, client.Age, client.Gender, client.Login, client.Password, client.Remove)
	if err != nil {
		fmt.Println("Cant insert user to DB err is ", err)
		return err
	}
	return err
}

func AddNewCurrency(db *sql.DB, currency models.Currency)(err error) {
	_, err = db.Exec("INSERT into currency(Name) VALUES (($1))", currency.Name)
	if err != nil {
		fmt.Println("Cant insert currency to DB err is ", err)
		return err
	}
	return err
}
/*

func RemoveById(db *sql.DB, id int64)(err error)  {
	_, err = db.Exec("UPDATE user SET Remove = true WHERE ID = ($1)", id)
	if err != nil {
		fmt.Println("cant update remove status, err is", err)
		return err
	}
	return err
}*/

