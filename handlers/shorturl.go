package handlers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"shorturl/data"
	"shorturl/db"

	"github.com/gorilla/mux"
)

type Shorturl struct {
	conn *db.DBConnection
}

func NewShorturl(c *db.DBConnection) *Shorturl {
	return &Shorturl{c}
}

func (s *Shorturl) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Printf("Service started..")
	switch r.Method {
	case http.MethodGet:
		s.getRealurl(rw, r)
	case http.MethodPost:
		s.addShorturl(rw, r)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (s *Shorturl) getRealurl(rw http.ResponseWriter, r *http.Request) {
	var surl string
	var err error
	vars := mux.Vars(r)
	surl = vars["url"]

	log.Println("surl:", surl)
	realUrl, err := data.GetRealurl(s.conn, surl)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(rw, "link not found", http.StatusNotFound)
			return
		}
		http.Error(rw, "Unable to fetch url details", http.StatusInternalServerError)
		return
	}
	http.Redirect(rw, r, realUrl, http.StatusFound)

}

func (s *Shorturl) addShorturl(rw http.ResponseWriter, r *http.Request) {
	var url data.Url
	var err error

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if string(b) == "{}" {
		http.Error(rw, "Empty request", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(b, &url); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	if url.UrlString == "" {
		http.Error(rw, "Should provide real url", http.StatusBadRequest)
		return
	}
	url, err = data.AddShorturl(s.conn, url.UrlString)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("url...", url)
	err = url.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
