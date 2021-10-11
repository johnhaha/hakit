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
