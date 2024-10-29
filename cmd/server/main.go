package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/antoniofmoraes/stress-test/cmd/server/status_picker"
)

func main() {
	httpStatusPicker, err := status_picker.NewHttpStatusPicker()
	if err != nil {
		log.Fatalf("Error at starting http status picker: %s", err)
	}

	http.NewServeMux()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		status := httpStatusPicker.PickRandom()
		w.WriteHeader(status)
		if status != 200 {
			fmt.Fprintf(w, "Error: %v", status)
		} else {
			w.Write([]byte("Hello, World!"))
		}
		log.Print(status)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
