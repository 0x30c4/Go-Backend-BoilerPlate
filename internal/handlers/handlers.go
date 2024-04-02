package handlers

import (
	"fmt"
	"net/http"
	"github.com/0x30c4/Go-Backend-BoilerPlate/internal/env"
)

func Index(w http.ResponseWriter, r *http.Request){
  userAgent := r.Header.Get("User-Agent")
  w.Header().Set("Content-Type", "text/html; charset=utf-8")

  if CheckIfBrowser(userAgent) {
    fmt.Fprintf(w, "You are a browser %s\n", env.DOMAIN)
  }else {
    fmt.Fprintf(w, "Not a browser %s\n", env.DOMAIN)
  }
}
