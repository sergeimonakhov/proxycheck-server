package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

//Proxy table
type Proxy struct {
	ID      int `json:"id"`
	ip      string
	port    string
	Country string `json:"country"`
	Respone int    `json:"respone"`
	Status  bool   `json:"status"`
	IPPort  string `json:"proxy"`
}

//Country list
type Country struct {
	id      int
	ip      string
	port    string
	Country string `json:"country"`
	respone int
	status  bool
}

//AllProxyReq get
func AllProxyReq(db *sql.DB) ([]*Proxy, error) {
	//rows, err = db.Query("SELECT id,ipport FROM proxy WHERE country = $1 ORDER BY respone", str)
	rows, err := db.Query("SELECT *, concat_ws(':', ip::inet, port::int) AS ipport FROM proxy")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*Proxy, 0)
	for rows.Next() {
		bk := new(Proxy)
		err = rows.Scan(&bk.ID, &bk.ip, &bk.port, &bk.Country, &bk.Respone, &bk.Status, &bk.IPPort)
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

//AllCountryReq get
func AllCountryReq(db *sql.DB) ([]*Country, error) {
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
	err := db.QueryRow("SELECT ipport FROM proxy WHERE ipport LIKE $1", ip+"%").Scan(&i)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		fmt.Println(err.Error())
	}
	return true
}

//AddToBase ip
func AddToBase(db *sql.DB, proxy string, country string, respone time.Duration, status string) {
	stmt, err := db.Prepare("INSERT INTO proxy VALUES (default, $1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(proxy, country, respone, status)
	if err != nil {
		log.Fatal(err)
	}
}

//InfoID get
/*func InfoID(db *sql.DB, id string) (ID, error) {
	var bks ID
	row := db.QueryRow("SELECT status FROM proxy WHERE id = $1", id)
	return bks, row.Scan(&bks.Status)
}*/
