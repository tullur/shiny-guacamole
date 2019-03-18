package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		log.Fatalln(err)
	}

	decode := schema.NewDecoder()
	if err := decode.Decode(dst, r.PostForm); err != nil {
		log.Fatalln(err)
	}

	return nil
}
