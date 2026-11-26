// Command backend — точка входа бэкенда. Пока только проверка, что сервер жив.
package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("ok"))
	})
	log.Printf("backend listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
