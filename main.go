package main

import (
	"fmt"
	"log"
	"net/http"
)

func tambah(a int, b int) int {
	return a + b
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Halo dari Go! Hasil: %d\n", tambah(10, 20))
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

	log.Println("Server berjalan di port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
