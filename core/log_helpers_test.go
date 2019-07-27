package core

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLogIncomingRequest(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	assert := assert.New(t)
	log.SetLevel(log.DebugLevel)
	requestWithBody, _ := http.NewRequest("POST", "testurl", bytes.NewBuffer([]byte(`{"hello": "world"}`)))
	requestWithBody.Header.Add("foo", "bar")

	LogIncomingRequest(requestWithBody)

	logOutput := buf.String()
	timeEndQuoteIndex := strings.IndexRune(logOutput[6:], '"') + 6
	logTime, _ := time.Parse(time.RFC3339, logOutput[6:timeEndQuoteIndex])
	parsedBody, _ := ioutil.ReadAll(requestWithBody.Body)

	assert.Equal("time=\"", logOutput[:6])
	assert.WithinDuration(time.Now(), logTime, time.Second)
	assert.Equal("\" level=debug msg=\"request received\" body=\"{\\\"hello\\\": \\\"world\\\"}\" header=\"map[Foo:[bar]]\" method=POST url=testurl\n", logOutput[timeEndQuoteIndex:])
	assert.Equal(`{"hello": "world"}`, string(parsedBody))
}

func TestLogIncomingRequest__WithoutBody(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	assert := assert.New(t)
	log.SetLevel(log.DebugLevel)
	requestWithBody, _ := http.NewRequest("POST", "testurl", bytes.NewBuffer([]byte{}))
	requestWithBody.Header.Add("foo", "bar")

	LogIncomingRequest(requestWithBody)

	logOutput := buf.String()
	timeEndQuoteIndex := strings.IndexRune(logOutput[6:], '"') + 6
	logTime, _ := time.Parse(time.RFC3339, logOutput[6:timeEndQuoteIndex])
	parsedBody, _ := ioutil.ReadAll(requestWithBody.Body)

	assert.Equal("time=\"", logOutput[:6])
	assert.WithinDuration(time.Now(), logTime, time.Second)
	assert.Equal("\" level=debug msg=\"request received\" header=\"map[Foo:[bar]]\" method=POST url=testurl\n", logOutput[timeEndQuoteIndex:])
	assert.Equal("", string(parsedBody))
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func TestLogIncomingRequest__WithErrorBody(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	assert := assert.New(t)
	log.SetLevel(log.DebugLevel)
	requestWithBody, _ := http.NewRequest("POST", "testurl", errReader(0))
	requestWithBody.Header.Add("foo", "bar")

	LogIncomingRequest(requestWithBody)

	logOutput := buf.String()
	timeEndQuoteIndex := strings.IndexRune(logOutput[6:], '"') + 6
	logTime, _ := time.Parse(time.RFC3339, logOutput[6:timeEndQuoteIndex])

	assert.Equal("time=\"", logOutput[:6])
	assert.WithinDuration(time.Now(), logTime, time.Second)
	assert.Equal("\" level=error msg=\"test error\" message=\"test error\" stack=\"[{github.com/ulodging/core/core/log_helpers_test.go:70: test error}]\"\n", logOutput[timeEndQuoteIndex:])
}
func TestLogErrorWithMessageAndStack(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	assert := assert.New(t)
	log.SetLevel(log.ErrorLevel)

	err := errors.New("test error")

	LogErrorWithMessageAndStack(err)

	logOutput := buf.String()
	timeEndQuoteIndex := strings.IndexRune(logOutput[6:], '"') + 6
	logTime, _ := time.Parse(time.RFC3339, logOutput[6:timeEndQuoteIndex])

	assert.Equal("time=\"", logOutput[:6])
	assert.WithinDuration(time.Now(), logTime, time.Second)
	assert.Equal("\" level=error msg=\"test error\" message=\"test error\" stack=\"[{github.com/ulodging/core/core/log_helpers_test.go:105: test error}]\"\n", logOutput[timeEndQuoteIndex:])
}

func TestLogWarningWithMessageAndStack(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	assert := assert.New(t)
	log.SetLevel(log.WarnLevel)

	err := errors.New("test error")

	LogWarningWithMessageAndStack(err)

	logOutput := buf.String()
	timeEndQuoteIndex := strings.IndexRune(logOutput[6:], '"') + 6
	logTime, _ := time.Parse(time.RFC3339, logOutput[6:timeEndQuoteIndex])

	assert.Equal("time=\"", logOutput[:6])
	assert.WithinDuration(time.Now(), logTime, time.Second)
	assert.Equal("\" level=warning msg=\"test error\" message=\"test error\" stack=\"[{github.com/ulodging/core/core/log_helpers_test.go:128: test error}]\"\n", logOutput[timeEndQuoteIndex:])
}
