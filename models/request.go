package models

import (
	"database/sql"
	"fmt"
)

//Proxy table
type Proxy struct {
	ID        int `json:"id"`
	ip        string
	port      string
	CountryID int     `json:"country"`
	Respone   float64 `json:"respone"`
	Status    bool    `json:"status"`
	IPPort    string  `json:"proxy"`
}

//Country table
type Country struct {
	ID      int    `json:"id"`
	Country string `json:"country"`
}

//ProxyStatus postreq
type ProxyStatus struct {
	Status bool
}

//AllProxyReq SELECT
func AllProxyReq(db *sql.DB) ([]*Proxy, error) {
	rows, err := db.Query("SELECT *, concat_ws(':', ip::inet, port::int) AS ipport FROM proxy")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*Proxy, 0)
	for rows.Next() {
		bk := new(Proxy)
		err = rows.Scan(&bk.ID, &bk.ip, &bk.port, &bk.CountryID, &bk.Respone, &bk.Status, &bk.IPPort)
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
		err = rows.Scan(&bk.ID, &bk.ip, &bk.port, &bk.CountryID, &bk.Respone, &bk.Status, &bk.IPPort)
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
	row, err := db.Query("SELECT *, concat_ws(':', ip::inet, port::int) AS ipport FROM proxy WHERE proxy_id = $1", id)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	bks := make([]*Proxy, 0)
	for row.Next() {
		bk := new(Proxy)
		err = row.Scan(&bk.ID, &bk.ip, &bk.port, &bk.CountryID, &bk.Respone, &bk.Status, &bk.IPPort)
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
	err := db.QueryRow("SELECT ip FROM proxy WHERE ip = $1", ip).Scan(&i)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		fmt.Println(err.Error())
	}
	return true
}

//AddToBase INSERT
func AddToBase(db *sql.DB, country string, ip string, port int, respone float64, status bool) {
	query := `WITH country_insert AS (
	   INSERT INTO country(country)
	   values($1)
	   ON CONFLICT (country) DO UPDATE
	   SET country = EXCLUDED.country
	   RETURNING country_id
	)
	  INSERT INTO proxy(ip, port, country_id, respone, status)
	  VALUES
	  ($2, $3, (SELECT country_id FROM country_insert), $4, $5)
	  ON CONFLICT (ip) DO NOTHING`

	stmt, err := db.Prepare(query)

	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = stmt.Exec(country, ip, port, respone, status)

	if err != nil {
		fmt.Println(err.Error())
	}
}

//UpdateStatus UPDATE
func UpdateStatus(db *sql.DB, id int, status bool) error {
	stmt, err := db.Prepare("UPDATE proxy SET status = $2 WHERE proxy_id = $1")

	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = stmt.Exec(id, status)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
