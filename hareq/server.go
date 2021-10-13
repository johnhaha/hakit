package hareq

import (
	"encoding/json"
	"net/http"
)

//http post req body parser
func BodyParser(r *http.Request, body interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(body)
	return err
}

//http send json
func SendJson(w http.ResponseWriter, data interface{}) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	return err
}

//http send string
func SendString(w http.ResponseWriter, data string) error {
	_, err := w.Write([]byte(data))
	return err
}

//http set header
func SetStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}
