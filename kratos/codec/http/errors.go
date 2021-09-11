package httpcodec

import (
	"encoding/json"
	"fmt"
	"net/http"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

type ErrorCodec struct {
	PlatformCode int32
	ServiceCode  int32
}

type ErrorCode int

type ErrorEncoder interface {
	Error() string
	GetCode() int
	GetPlatformCode() int
	GetServiceCode() int
	GetErrTypeCode() int
}

type ErrResp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func ErrorEncode(w http.ResponseWriter, r *http.Request, err error) {
	m, ok := err.(ErrorEncoder)
	if !ok {
		khttp.DefaultErrorEncoder(w, r, err)
	}
	data, err := json.Marshal(&ErrResp{
		Code: fmt.Sprintf("%v%v%v%v", m.GetPlatformCode(), m.GetServiceCode(), m.GetErrTypeCode(), m.GetCode()),
		Msg:  m.Error(),
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
