package core

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"
)

func LogIncomingRequest(r *http.Request) {
	// Parse request body
	parsedBody, parseErr := ioutil.ReadAll(r.Body)

	// If request body parses without issue and has content, log body
	// if request body parses without issue but has no content, log the rest
	if parseErr == nil && len(parsedBody) > 0 {
		log.WithFields(log.Fields{
			"method": r.Method,
			"url":    r.URL,
			"header": r.Header,
			"body":   string(parsedBody),
		}).Debug("request received")
	} else if parseErr == io.EOF || (len(parsedBody) == 0 && parseErr == nil) {
		log.WithFields(log.Fields{
			"method": r.Method,
			"header": r.Header,
			"url":    r.URL,
		}).Debug("request received")
	} else {
		LogErrorWithMessageAndStack(parseErr)
	}
	// repopulate the request body
	r.Body = ioutil.NopCloser(bytes.NewBuffer(parsedBody))
}

func LogErrorWithMessageAndStack(err error) {
	log.WithFields(log.Fields{
		"message": err.Error(),
		"stack":   errors.Details(err),
	}).Error(err.Error())
}
