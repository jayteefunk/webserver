package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"log"
)

func main() {
	http.HandleFunc("/", myhandler)

	http.ListenAndServe(":8080", nil)
}
func myhandler(w http.ResponseWriter, r *http.Request) {
	var file *os.File
	log.Print(r.URL.Path)
	var err error
	if r.URL.Path == "/" {
		file, err = os.Open("index.html")
	} else {
		path := []byte(r.URL.Path)[1:]
		file, err = os.Open(string(path))
		w.Header().Set("Content-Type", "text/css")
	defer file.Close()
	}
	if err != nil {
		fmt.Print(err)
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Print(err)
	}
	w.Write(fileBytes)
}
