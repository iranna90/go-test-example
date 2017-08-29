package main

import (
	"net/http"
	"io"
	"os"
	"log"
	"fmt"
)

var mux = serverMux

func serverMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/root", Handler)
	return mux
}

func main() {
	log.Println("starting the server")
	http.ListenAndServe(":9876", mux())
}

var reader = filReader

func Handler(w http.ResponseWriter, _ *http.Request) {
	dataBytes := make([]byte, 256)
	bytes, err := reader().Read(dataBytes)

	if err != nil {
		message := "error while reading data from reader"
		log.Println(message, err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}
	log.Println(fmt.Sprintf("Read %d number of bytes from reader", bytes))
	var exactData []byte
	for i := 0; i < bytes; i++ {
		exactData = append(exactData, dataBytes[i])
	}
	count, err := w.Write(exactData)
	if err != nil {
		message := "error while writing data to response"
		log.Println(message, err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}
	log.Println(fmt.Sprintf("%d number of bytes written successfully", count))
}

func filReader() io.Reader {
	fileName := "/Users/iranna.patil/go/check.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
