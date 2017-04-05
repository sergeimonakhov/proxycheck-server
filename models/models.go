package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

//ListProxy get
func ListProxy(db *sql.DB, str string) ([]*Proxy, error) {
	var (
		rows *sql.Rows
		err  error
	)

	if str != "" {
		rows, err = db.Query("SELECT ipport FROM proxy WHERE country = $1", str)
	} else {
		rows, err = db.Query("SELECT ipport FROM proxy")
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*Proxy, 0)
	for rows.Next() {
		bk := new(Proxy)
		err = rows.Scan(&bk.Proxy)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}

//ListCountry get
func ListCountry(db *sql.DB) ([]*Country, error) {

	rows, err := db.Query("SELECT country FROM proxy GROUP BY country")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*Country, 0)
	for rows.Next() {
		bk := new(Country)
		err = rows.Scan(&bk.Country)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}

//ExistIP test
func ExistIP(db *sql.DB, ip string) bool {
	var i string
	err := db.QueryRow("SELECT ipport FROM proxy WHERE ipport LIKE $1 LIMIT 1", ip+"%").Scan(&i)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		fmt.Println(err.Error())
	}
	return true
}

//AddToBase ip
func AddToBase(db *sql.DB, proxy string, country string, respone time.Duration) {
	stmt, err := db.Prepare("INSERT INTO proxy VALUES (default, $1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(proxy, country, respone)
	if err != nil {
		log.Fatal(err)
	}
}
