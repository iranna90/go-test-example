package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io"
	"strings"
)

var str = "from string reader"

func TestRestCall(t *testing.T) {
	server := httptest.NewServer(nil)
	defer server.Close()
	http.Head(server.URL)
}

func HandleTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func TestRecordedOutPut(t *testing.T) {
	request, _ := http.NewRequest("GET", "http://localhost:1111/root", nil)

	writer := httptest.NewRecorder()

	Handler(writer, request)
	responseData := writer.Body.String()
	// default data is from file
	fileData := "hello welcome to go lang"
	if responseData != fileData {
		t.Errorf("Data did not read is %s and expected from file is %s", fileData, responseData)
	}

	// provide reader for string
	reader = stringReader
	writer = httptest.NewRecorder()
	Handler(writer, request)
	responseData = writer.Body.String()
	if responseData != str {
		t.Errorf("Data retrived is %s and data from string reader %s", responseData, str)
	}
}

func stringReader() io.Reader {
	return strings.NewReader(str)
}