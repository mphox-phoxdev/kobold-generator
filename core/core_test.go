package core

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/juju/errors"
	"github.com/stretchr/testify/assert"
)

func TestCheckForRequestBody(t *testing.T) {
	requestWithBody, _ := http.NewRequest("POST", "testurl", bytes.NewBufferString("test body"))
	assert := assert.New(t)
	err := CheckForRequestBody(requestWithBody)
	assert.Nil(err, "there shold be no errors thrown")
}
func TestCheckForRequestBody_EmptyRequest(t *testing.T) {
	emptyRequest, _ := http.NewRequest("POST", "testurl", nil)
	assert := assert.New(t)
	err := CheckForRequestBody(emptyRequest)
	assert.NotNil(err, "there shold be an error thrown")
	assert.Equal(true, errors.IsNotValid(err), "the error should be of type NotValid")
}

func TestWriteJSONResponse(t *testing.T) {
	writer := httptest.NewRecorder()
	assert := assert.New(t)

	WriteJSONResponse(writer, struct{ Hello string }{Hello: "World"})
	resp := writer.Result()

	assert.Equal(http.StatusOK, resp.StatusCode, "the response code should be 200")
	assert.Equal("application/json", resp.Header.Get("Content-Type"), "the content type should be \"application/json\"")
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("{\"Hello\":\"World\"}\n", string(body), "the body should be a serialized version of the input struct")
}

func TestWriteStatusOKResponse(t *testing.T) {
	writer := httptest.NewRecorder()
	assert := assert.New(t)

	WriteStatusOKResponse(writer)
	resp := writer.Result()

	assert.Equal(http.StatusOK, resp.StatusCode, "the response code should be 200")
	assert.Equal("application/json", resp.Header.Get("Content-Type"), "the content type should be \"application/json\"")
	body, _ := ioutil.ReadAll(resp.Body)

	var r response
	_ = json.Unmarshal(body, &r)
	assert.Equal(http.StatusOK, r.Status, "the response status should be 200")
	assert.Equal(http.StatusText(http.StatusOK), r.Message)
}

func TestWriteStatusOKResponseWithMessage(t *testing.T) {
	writer := httptest.NewRecorder()
	assert := assert.New(t)

	WriteStatusOKResponseWithMessage(writer, "boop")
	resp := writer.Result()

	assert.Equal(http.StatusOK, resp.StatusCode, "the response code should be 200")
	assert.Equal("application/json", resp.Header.Get("Content-Type"), "the content type should be \"application/json\"")
	body, _ := ioutil.ReadAll(resp.Body)

	var r response
	_ = json.Unmarshal(body, &r)
	assert.Equal(http.StatusOK, r.Status, "the response status should be 200")
	assert.Equal("boop", r.Message)
}
func TestWriteInternalServerErrorResponse(t *testing.T) {
	writer := httptest.NewRecorder()
	assert := assert.New(t)

	WriteInternalServerErrorResponse(writer)
	resp := writer.Result()

	assert.Equal(http.StatusInternalServerError, resp.StatusCode, "the response code should be 500")
	assert.Equal("application/json", resp.Header.Get("Content-Type"), "the content type should be \"application/json\"")
	body, _ := ioutil.ReadAll(resp.Body)

	var r response
	_ = json.Unmarshal(body, &r)
	assert.Equal(http.StatusInternalServerError, r.Status, "the response status should be 500")
	assert.Equal(http.StatusText(http.StatusInternalServerError), r.Message)
}

func TestWriteInternalServerErrorResponseWithMessage(t *testing.T) {
	writer := httptest.NewRecorder()
	assert := assert.New(t)

	WriteInternalServerErrorResponseWithMessage(writer, "boop")
	resp := writer.Result()

	assert.Equal(http.StatusInternalServerError, resp.StatusCode, "the response code should be 500")
	assert.Equal("application/json", resp.Header.Get("Content-Type"), "the content type should be \"application/json\"")
	body, _ := ioutil.ReadAll(resp.Body)

	var r response
	_ = json.Unmarshal(body, &r)
	assert.Equal(http.StatusInternalServerError, r.Status, "the response status should be 500")
	assert.Equal("boop", r.Message)
}
func TestWriteBadRequestErrorResponse(t *testing.T) {
	writer := httptest.NewRecorder()
	assert := assert.New(t)

	WriteBadRequestErrorResponse(writer)
	resp := writer.Result()

	assert.Equal(http.StatusBadRequest, resp.StatusCode, "the response code should be 400")
	assert.Equal("application/json", resp.Header.Get("Content-Type"), "the content type should be \"application/json\"")
	body, _ := ioutil.ReadAll(resp.Body)

	var r response
	_ = json.Unmarshal(body, &r)
	assert.Equal(http.StatusBadRequest, r.Status, "the response status should be 400")
	assert.Equal(http.StatusText(http.StatusBadRequest), r.Message)
}

func TestWriteBadRequestErrorResponseWithMessage(t *testing.T) {
	writer := httptest.NewRecorder()
	assert := assert.New(t)

	WriteBadRequestErrorResponseWithMessage(writer, "hello world")
	resp := writer.Result()

	assert.Equal(http.StatusBadRequest, resp.StatusCode, "the response code should be 400")
	assert.Equal("application/json", resp.Header.Get("Content-Type"), "the content type should be \"application/json\"")
	body, _ := ioutil.ReadAll(resp.Body)

	var r response
	_ = json.Unmarshal(body, &r)
	assert.Equal(http.StatusBadRequest, r.Status, "the response status should be 400")
	assert.Equal("hello world", r.Message)
}

func TestWriteNotFoundErrorResponse(t *testing.T) {
	writer := httptest.NewRecorder()
	assert := assert.New(t)

	WriteNotFoundErrorResponse(writer)
	resp := writer.Result()

	assert.Equal(http.StatusNotFound, resp.StatusCode, "the response code should be 404")
	assert.Equal("application/json", resp.Header.Get("Content-Type"), "the content type should be \"application/json\"")
	body, _ := ioutil.ReadAll(resp.Body)

	var r response
	_ = json.Unmarshal(body, &r)
	assert.Equal(http.StatusNotFound, r.Status, "the response status should be 404")
	assert.Equal(http.StatusText(http.StatusNotFound), r.Message)
}

func TestWriteNotFoundErrorResponseWithMessage(t *testing.T) {
	writer := httptest.NewRecorder()
	assert := assert.New(t)

	WriteNotFoundErrorResponseWithMessage(writer, "hello world")
	resp := writer.Result()

	assert.Equal(http.StatusNotFound, resp.StatusCode, "the response code should be 404")
	assert.Equal("application/json", resp.Header.Get("Content-Type"), "the content type should be \"application/json\"")
	body, _ := ioutil.ReadAll(resp.Body)

	var r response
	_ = json.Unmarshal(body, &r)
	assert.Equal(http.StatusNotFound, r.Status, "the response status should be 404")
	assert.Equal("hello world", r.Message)
}

func TestWriteForbiddenErrorResponse(t *testing.T) {
	writer := httptest.NewRecorder()
	assert := assert.New(t)

	WriteForbiddenErrorResponse(writer)
	resp := writer.Result()

	assert.Equal(http.StatusForbidden, resp.StatusCode, "the response code should be 403")
	assert.Equal("application/json", resp.Header.Get("Content-Type"), "the content type should be \"application/json\"")
	body, _ := ioutil.ReadAll(resp.Body)

	var r response
	_ = json.Unmarshal(body, &r)
	assert.Equal(http.StatusForbidden, r.Status, "the response status should be 403")
	assert.Equal(http.StatusText(http.StatusForbidden), r.Message)
}

func TestWriteForbiddenErrorResponseWithMessage(t *testing.T) {
	writer := httptest.NewRecorder()
	assert := assert.New(t)

	WriteForbiddenErrorResponseWithMessage(writer, "boop")
	resp := writer.Result()

	assert.Equal(http.StatusForbidden, resp.StatusCode, "the response code should be 403")
	assert.Equal("application/json", resp.Header.Get("Content-Type"), "the content type should be \"application/json\"")
	body, _ := ioutil.ReadAll(resp.Body)

	var r response
	_ = json.Unmarshal(body, &r)
	assert.Equal(http.StatusForbidden, r.Status, "the response status should be 403")
	assert.Equal("boop", r.Message)
}

func TestWriteUnauthorizedErrorResponse(t *testing.T) {
	writer := httptest.NewRecorder()
	assert := assert.New(t)

	WriteUnauthorizedErrorResponse(writer)
	resp := writer.Result()

	assert.Equal(http.StatusUnauthorized, resp.StatusCode, "the response code should be 401")
	assert.Equal("application/json", resp.Header.Get("Content-Type"), "the content type should be \"application/json\"")
	body, _ := ioutil.ReadAll(resp.Body)

	var r response
	_ = json.Unmarshal(body, &r)
	assert.Equal(http.StatusUnauthorized, r.Status, "the response status should be 401")
	assert.Equal(http.StatusText(http.StatusUnauthorized), r.Message)
}

func TestWriteUnauthorizedErrorResponseWithMessage(t *testing.T) {
	writer := httptest.NewRecorder()
	assert := assert.New(t)

	WriteUnauthorizedErrorResponseWithMessage(writer, "boop")
	resp := writer.Result()

	assert.Equal(http.StatusUnauthorized, resp.StatusCode, "the response code should be 401")
	assert.Equal("application/json", resp.Header.Get("Content-Type"), "the content type should be \"application/json\"")
	body, _ := ioutil.ReadAll(resp.Body)

	var r response
	_ = json.Unmarshal(body, &r)
	assert.Equal(http.StatusUnauthorized, r.Status, "the response status should be 401")
	assert.Equal("boop", r.Message)
}

func TestWriteResponse(t *testing.T) {
	writer := httptest.NewRecorder()
	assert := assert.New(t)

	writeResponse(writer, 400)
	resp := writer.Result()

	assert.Equal(400, resp.StatusCode, "the response code should be 400")
	assert.Equal("application/json", resp.Header.Get("Content-Type"), "the content type should be \"application/json\"")
	body, _ := ioutil.ReadAll(resp.Body)

	var r response
	_ = json.Unmarshal(body, &r)
	assert.Equal(400, r.Status, "the response status should be 400")
	assert.Equal(http.StatusText(400), r.Message, "the response body should be \"hello world\"")
}

func TestWriteResponseWithMessage(t *testing.T) {
	writer := httptest.NewRecorder()
	assert := assert.New(t)

	writeResponseWithMessage(writer, 123, "hello world")
	resp := writer.Result()

	assert.Equal(123, resp.StatusCode, "the response code should be 123")
	assert.Equal("application/json", resp.Header.Get("Content-Type"), "the content type should be \"application/json\"")
	body, _ := ioutil.ReadAll(resp.Body)

	var r response
	_ = json.Unmarshal(body, &r)
	assert.Equal(123, r.Status, "the response status should be 123")
	assert.Equal("hello world", r.Message, "the response body should be \"hello world\"")
}
