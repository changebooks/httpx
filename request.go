package httpx

import (
	"github.com/changebooks/http"
	"github.com/changebooks/log"
	http2 "net/http"
	"time"
)

func (x *Http) Request(idRegister *log.IdRegister, req *http2.Request, timeout time.Duration) *http.Schema {
	tag := "Request"

	start := time.Now()

	r := x.talk.Request(req, timeout)

	done := time.Now()
	remark := NewRemark(r, start, done, timeout)

	if r.ReadError != nil {
		x.logger.E(tag, ReadFailure, remark, r.ReadError, "", idRegister)
	}

	if r.Error == nil {
		x.logger.I(tag, Success, remark, idRegister)
	} else {
		x.logger.E(tag, Failure, remark, r.Error, "", idRegister)
	}

	return r
}
