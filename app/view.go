package app

import (
	"encoding/json"
	"log"
	"net/http"
)

func Render(w http.ResponseWriter, code int, obj interface{}) {
	enc, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(enc); err != nil {
		log.Fatal(err)
		return
	}
}
