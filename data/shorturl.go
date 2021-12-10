package data

import (
	"crypto/sha1"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"shorturl/db"
)

const table = "gotool.Shorturl"
const baseurl = "http://localhost:8080/"

type Url struct {
	UrlString string `json:"url"`
}

func (surl *Url) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(surl)
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
func AddShorturl(con *db.DBConnection, url string) (surl Url, err error) {
	var sqlRows *sql.Rows
	var short_url string

	if short_url, err = getShorturl(con, url); err == sql.ErrNoRows {
		ib := con.Builder.NewInsertBuilder()
		url_hash := sha1.Sum([]byte(url))

		short_url = fmt.Sprintf("%x", url_hash[0:4])
		log.Println("short url...", short_url)

		ib.InsertInto(table).Cols("long_url", "short_url").Values(url, short_url)

		sqlStr, args := ib.Build()
		log.Println(sqlStr, args)
		if sqlRows, err = con.DB.Query(sqlStr, args...); err != nil {
			return
		}

		defer sqlRows.Close()

	}

	surl = Url{
		UrlString: baseurl + short_url,
	}
	return
}

func getShorturl(con *db.DBConnection, url string) (short_url string, err error) {
	var sqlRows *sql.Rows

	sb := con.Builder.NewSelectBuilder()
	sb.Select("short_url").From(table).Where(sb.Equal("long_url", url))
	sqlStr, args := sb.Build()
	log.Printf("query : %s", sqlStr)

	if sqlRows, err = con.DB.Query(sqlStr, args...); err != nil {
		return
	}

	defer sqlRows.Close()
	for sqlRows.Next() {
		if err = sqlRows.Scan(&short_url); err != nil {
			return
		}
	}
	if short_url == "" {
		return short_url, sql.ErrNoRows
	}
	return
}
