package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"example.com/cbench/handler-nethttp/liba"
)

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/fingerprint", fingerprint)
	if len(os.Args) > 1 && os.Args[1] == "serve" {
		if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}
	// Without "serve", answer one request in process and exit.
	rec := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/fingerprint?data=hello", nil))
	fmt.Println(rec.Code, rec.Body.String())
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func fingerprint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%x", liba.Fingerprint([]byte(r.URL.Query().Get("data"))))
}
