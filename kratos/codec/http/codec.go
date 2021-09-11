package errors

import (
	"encoding/json"
	"net/http"
)

const (
	baseContentType = "application/json"
	StatusOK        = 200
)

// ResponseEncoder encodes the object to the HTTP response.
func ResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	type response struct {
		Code int32
		Data interface{}
	}
	data, err := json.Marshal(response{
		Code: StatusOK,
		Data: v,
	})
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", baseContentType)
	_, err = w.Write(data)

	if err != nil {
		return err
	}
	return nil
}
