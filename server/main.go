package main

import (
	"net/http"
	"os"

	"github.com/caarlos0/getantibody"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		release, err := getantibody.LatestRelease()
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.Write([]byte(release))
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, nil)
}
