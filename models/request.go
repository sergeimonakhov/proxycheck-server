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

//Country table
type Country struct {
	ID      int    `json:"id"`
	Country string `json:"country"`
}

//AllProxyReq SELECT
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

//AllCountryReq SELECT
func AllCountryReq(db *sql.DB) ([]*Country, error) {
	rows, err := db.Query("SELECT * FROM country")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*Country, 0)
	for rows.Next() {
		bk := new(Country)
		err = rows.Scan(&bk.ID, &bk.Country)
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

//FilterCountryReq SELECT
func FilterCountryReq(db *sql.DB, id int) ([]*Proxy, error) {
	rows, err := db.Query("SELECT *, concat_ws(':', ip::inet, port::int) AS ipport FROM proxy WHERE country_id = $1", id)

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

//FilterProxyReq SELECT
func FilterProxyReq(db *sql.DB, id int) ([]*Proxy, error) {
	row, err := db.Query("SELECT *, concat_ws(':', ip::inet, port::int) AS ipport FROM proxy WHERE id_proxy = $1", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	bks := make([]*Proxy, 0)
	for row.Next() {
		bk := new(Proxy)
		err = row.Scan(&bk.ID, &bk.ip, &bk.port, &bk.Country, &bk.Respone, &bk.Status, &bk.IPPort)
		if err != nil {
			return nil, err
		}

		bks = append(bks, bk)
	}
	if err = row.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}

//ExistIP test
func ExistIP(db *sql.DB, ip string) bool {
	var i string
	err := db.QueryRow("SELECT * FROM proxy WHERE ip = $1", ip).Scan(&i)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		fmt.Println(err.Error())
	}
	return true
}

//AddToBase INSERT
func AddToBase(db *sql.DB, country string, ip string, port int, respone time.Duration, status bool) {
	query := `WITH country_insert AS (
	   INSERT INTO country(country)
	   values($1)
	   ON CONFLICT (country) DO UPDATE
	   SET country = excluded.country
	   RETURNING id_country
	)
	  INSERT INTO proxy(ip, port, country_id, respone, status)
	  VALUES
	  ($2, $3, (SELECT id_country FROM country_insert), $4, $5)
	  ON CONFLICT (ip) DO NOTHING`

	stmt, err := db.Prepare(query)

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(country, ip, port, respone, status)

	if err != nil {
		log.Fatal(err)
	}
}
