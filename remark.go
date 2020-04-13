package httpx

import (
	"github.com/changebooks/http"
	http2 "net/http"
	"net/url"
	"time"
)

type Remark struct {
	Scheme   string        `json:"scheme"`
	Proto    string        `json:"proto"`
	Host     string        `json:"host"`
	Path     string        `json:"path"`
	Method   string        `json:"method"`
	Header   http2.Header  `json:"header"`
	Form     url.Values    `json:"form"`
	PostForm url.Values    `json:"postForm"`
	Timeout  time.Duration `json:"timeout"`
	Retries  int           `json:"retries"`
	Status   string        `json:"status"`
	Total    time.Duration `json:"total"`
	Start    time.Time     `json:"start"`
	Done     time.Time     `json:"done"`
}

func NewRemark(schema *http.Schema, start time.Time, done time.Time, timeout time.Duration) *Remark {
	total := done.Sub(start)

	r := &Remark{
		Timeout: timeout,
		Total:   total,
		Start:   start,
		Done:    done,
	}

	if schema == nil {
		return r
	}

	r.Retries = schema.Retries

	req := schema.Request
	if req != nil {
		r.Proto = req.Proto
		r.Host = req.Host
		r.Method = req.Method
		r.Header = req.Header
		r.Form = req.Form
		r.PostForm = req.PostForm

		if req.URL != nil {
			r.Scheme = req.URL.Scheme
			r.Path = req.URL.Path
		}
	}

	resp := schema.Response
	if resp != nil {
		r.Status = resp.Status
	}

	return r
}

type HttpRequestRemark struct {
	Method string            `json:"method"`
	Path   string            `json:"path"`
	Params map[string]string `json:"params"`
	Url    string            `json:"url"`
	Scheme string            `json:"scheme"`
	Proto  string            `json:"proto"`
	Host   string            `json:"host"`
}

func NewHttpRequestRemark(method string, path string, params map[string]string, req *http2.Request, url string) *HttpRequestRemark {
	r := &HttpRequestRemark{
		Method: method,
		Path:   path,
		Params: params,
		Url:    url,
	}

	if req != nil {
		r.Proto = req.Proto
		r.Host = req.Host
		if req.URL != nil {
			r.Scheme = req.URL.Scheme
		}
	}

	return r
}
