package main

import (
	"net/http"

	"github.com/DylanTet/go-mindshift/internal/interfaces/router"
)

func main() {
  getUser := func(w http.ResponseWriter, _ *http.Request) {
  }

  http.HandleFunc("/get-user", getUser)
}
