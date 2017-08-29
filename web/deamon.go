package main

import (
	"net/http"
	"io"
	"os"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/root", Handler)
	log.Println("starting the server")
	http.ListenAndServe(":1111", nil)
}

var reader = filReader

func Handler(w http.ResponseWriter, _ *http.Request) {
	dataBytes := make([]byte, 256)
	bytes, err := reader().Read(dataBytes)

	if err != nil {
		message := "error while reading data from reader"
		log.Fatal(message, err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}
	log.Println(fmt.Sprintf("Read %d number of bytes from reader", bytes))
	var exactData []byte
	for i := 0; i < bytes; i++ {
		exactData = append(exactData,dataBytes[i])
	}
	count, err := w.Write(exactData)
	if err != nil {
		message := "error while writing data to response"
		log.Fatal(message, err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}
	log.Println(fmt.Sprintf("%d number of bytes written successfully", count))
}

func filReader() io.Reader {
	fileName := "/home/iranna/go/check.text"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
