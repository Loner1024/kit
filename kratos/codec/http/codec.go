package httpcodec

import (
	"net/http"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

const (
	baseContentType = "application/json"
	StatusOK        = 200
)

// ResponseEncoder encodes the object to the HTTP response.
func ResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	codec, _ := khttp.CodecForRequest(r, "Accept")
	type response struct {
		Code int32       `json:"code"`
		Data interface{} `json:"data"`
	}
	data, err := codec.Marshal(response{
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
