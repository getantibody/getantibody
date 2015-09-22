package main

import (
	"net/http"
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
	http.ListenAndServe(":3000", nil)
}
