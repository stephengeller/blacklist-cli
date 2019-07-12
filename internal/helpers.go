package internal

import "log"

type Logger interface {
	Fatal(err error)
}

type DefaultLogger struct {
}

func (logger *DefaultLogger) Fatal(e error) {
	log.Fatal(e)
}

func Check(logger Logger, err error) {
	if err != nil {
		logger.Fatal(err)
	}
}
