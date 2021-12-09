package data

import (
	"crypto/sha1"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"shorturl/db"
)

const table = "gotool.Shorturl"

type Shorturl struct {
	Id       int64  `json:"id"`
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}

func (sd *Shorturl) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(sd)
}

// GetRealurl return Shorturl for the real one
func GetRealurl(con *db.DBConnection, url string) (real_url string, err error) {
	var sqlRows *sql.Rows

	sb := con.Builder.NewSelectBuilder()
	sb.Select("long_url").From(table).Where(sb.Equal("short_url", url))
	sqlStr, args := sb.Build()
	log.Printf("query : %s", sqlStr)

	if sqlRows, err = con.DB.Query(sqlStr, args...); err != nil {
		return
	}

	defer sqlRows.Close()
	for sqlRows.Next() {
		if err = sqlRows.Scan(&real_url); err != nil {
			return
		}
	}
	if real_url == "" {
		return real_url, sql.ErrNoRows
	}
	return
}

// AddShorturl save Shorturl
func AddShorturl(con *db.DBConnection, url string) (short_url string, err error) {
	var sqlRows *sql.Rows

	ib := con.Builder.NewInsertBuilder()
	url_hash := sha1.Sum([]byte(url))
	short_url = string(url_hash[0:6])

	ib.InsertInto(table).Cols("long_url", "short_url").Values(url, short_url)

	sqlStr, args := ib.Build()
	log.Println(sqlStr, args)
	if sqlRows, err = con.DB.Query(sqlStr, args...); err != nil {
		return
	}

	defer sqlRows.Close()

	return
}
