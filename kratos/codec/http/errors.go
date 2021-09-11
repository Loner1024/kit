package httpcodec

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
)

var (
	PlatformCode int32
	ServiceCode  int32
)

type ErrResp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	se := errors.FromError(err)
	data, err := json.Marshal(&ErrResp{
		Code: fmt.Sprintf("%v%v%v", PlatformCode, ServiceCode, se.Code),
		Msg:  se.Message,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", baseContentType)
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
