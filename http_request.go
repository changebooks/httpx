package httpx

import (
	"github.com/changebooks/log"
	http2 "net/http"
)

func (x *Http) NewHttpRequest(idRegister *log.IdRegister,
	method string, path string, params map[string]string) (req *http2.Request, url string, err error) {
	tag := "NewHttpRequest"

	req, url, err = x.talk.NewRequest(method, path, params)

	remark := NewHttpRequestRemark(method, path, params, req, url)

	if err == nil {
		x.logger.I(tag, Success, remark, idRegister)
	} else {
		x.logger.E(tag, Failure, remark, err, "", idRegister)
	}

	return
}
