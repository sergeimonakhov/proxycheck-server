package models

import (
	"database/sql"
)

//ListProxy get
func ListProxy(db *sql.DB, str string) ([]*Proxy, error) {
	var (
		rows *sql.Rows
		err  error
	)

	if str != "" {
		rows, err = db.Query("SELECT * FROM proxys WHERE country = $1", str)
	} else {
		rows, err = db.Query("SELECT * FROM proxys")
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*Proxy, 0)
	for rows.Next() {
		bk := new(Proxy)
		err = rows.Scan(&bk.id, &bk.Proxy, &bk.Country, &bk.respone)
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
func ListCountry(db *sql.DB) ([]*Proxy, error) {

	rows, err := db.Query("SELECT country FROM proxys GROUP BY country")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*Proxy, 0)
	for rows.Next() {
		bk := new(Proxy)
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
