package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareFormHandlerSuccess(t *testing.T) {
	requestPath := "/ping"
	var handler = Handler{}

	var writer = httptest.NewRecorder()

	var expectedResponseCode = 200
	var expectedResponseBody = "\"pong\"\n"

	req, err := http.NewRequest("GET", requestPath, strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}

	//add the cookie
	handler.PingHandler(writer, req)

	assert.Equal(t, expectedResponseCode, writer.Code)
	assert.Equal(t, expectedResponseBody, writer.Body.String())
}
