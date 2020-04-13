package httpx

import (
	"errors"
	"github.com/changebooks/http"
	"github.com/changebooks/log"
)

type Http struct {
	talk   *http.Talk
	logger *log.Logger
}

func New(talk *http.Talk, logger *log.Logger) (*Http, error) {
	if talk == nil {
		return nil, errors.New("talk can't be nil")
	}

	if logger == nil {
		return nil, errors.New("logger can't be nil")
	}

	return &Http{
		talk:   talk,
		logger: logger,
	}, nil
}

func (x *Http) GetTalk() *http.Talk {
	return x.talk
}

func (x *Http) GetLogger() *log.Logger {
	return x.logger
}
