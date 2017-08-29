package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io"
	"strings"
	"io/ioutil"
	"fmt"
)

var str = "from string reader"

func TestRestCall(t *testing.T) {
	server := httptest.NewServer(mux())
	defer server.Close()
	// provide reader for string
	reader = stringReader

	response, err := http.Head(server.URL + "/root")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("status ", response.StatusCode)
	bytes, err := ioutil.ReadAll(response.Body)
	fmt.Println("response length ", len(bytes))
	fmt.Println("response ", string(bytes))

	// provide reader for string
	reader = filReader
	response, err = http.Get(server.URL + "/root")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("status ", response.StatusCode)
	bytes, err = ioutil.ReadAll(response.Body)
	fmt.Println("response length ", len(bytes))
	fmt.Println("response ", string(bytes))
}

func TestRecordedOutPut(t *testing.T) {
	request, _ := http.NewRequest("GET", "http://localhost:1111/root", nil)

	writer := httptest.NewRecorder()

	Handler(writer, request)
	responseData := writer.Body.String()
	// default data is from file
	fileData := "hello welcome to go lang"
	if responseData != fileData {
		t.Errorf("Data retrieved is \"%s\" and expected from file is \"%s\"", fileData, responseData)
	}

	// provide reader for string
	reader = stringReader
	writer = httptest.NewRecorder()
	Handler(writer, request)
	responseData = writer.Body.String()
	if responseData != str {
		t.Errorf("Data retrived is \"%s\" and data from string reader \"%s\"", responseData, str)
	}
}

func stringReader() io.Reader {
	return strings.NewReader(str)
}
