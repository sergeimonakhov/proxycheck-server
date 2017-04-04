package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//CountryProxy get
func CountryProxy(db *sql.DB, str string) ([]*Proxy, error) {
	var rows *sql.Rows
	var err error
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
