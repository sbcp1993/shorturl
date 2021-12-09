package handlers

import (
	"log"
	"net/http"

	"shorturl/db"
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
		// call get method
	case http.MethodPost:
		// call post method
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}

}
