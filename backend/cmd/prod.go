package cmd

import (
	"net/http"
)

func Prod() {
    http.Handle("/", http.FileServer(http.Dir("../../frontend/dist")))
    http.ListenAndServe(":3000", nil)
}
