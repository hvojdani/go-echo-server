package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := getenv("PORT", "8080")
	readTimeout := getenv("READ_TIMEOUT", "5s")
	writeTimeout := getenv("WRITE_TIMEOUT", "10s")

	rt, err := time.ParseDuration(readTimeout)
	if err != nil {
		log.Printf("invalid READ_TIMEOUT %q, using 5s", readTimeout)
		rt = 5 * time.Second
	}
	wt, err := time.ParseDuration(writeTimeout)
	if err != nil {
		log.Printf("invalid WRITE_TIMEOUT %q, using 10s", writeTimeout)
		wt = 10 * time.Second
	}

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      http.HandlerFunc(logHandler),
		ReadTimeout:  rt,
		WriteTimeout: wt,
	}

	log.Printf("starting server on :%s", port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading body: %v", err)
	}
	_ = r.Body.Close()

	logdata := fmt.Sprintf("%s %s body=%s", r.Method, r.URL.String(), string(body))
	log.Printf("%s", logdata)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, logdata)
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
