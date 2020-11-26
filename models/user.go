package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID int64
	Name string
	Surname string
	Age int64
	Gender string
	Login string
	Password string
	Remove bool
}

func (user User) AddNewAccount(db *sql.DB, account Account) (err error) {
		_, err = db.Exec("INSERT into account(UserID, Number, Amount, CurrencyID) VALUES (($1),($2),($3),($4))", user.ID, account.Number, account.Amount, account.CurrencyID)
		if err != nil {
			fmt.Println("Cant insert account to DB err is ", err)
			return err
		}
		return err

}
