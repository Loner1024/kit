package httpcodec

import (
	"net/http"

	"github.com/Loner1024/kit/kratos/codec/http/jsonsb"
)

const (
	baseContentType = "application/json"
	StatusOK        = 200
)

// ResponseEncoder encodes the object to the HTTP response.
func ResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	type response struct {
		Code int32       `json:"code"`
		Data interface{} `json:"data"`
	}
	data, err := jsonsb.Marshal(response{
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
