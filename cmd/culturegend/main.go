package main

import (
	"log"
  "fmt"
  "math/rand"
  "net/http"
  "time"

  "github.com/ironarachne/culturegen"
  "github.com/ironarachne/random"
  "github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func getCulture(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "id")

	var newCulture culturegen.Culture

	random.SeedFromString(id)

	newCulture = culturegen.GenerateCulture()

	json.NewEncoder(w).Encode(newCulture)
}

func main() {
  r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Use(middleware.Timeout(60 * time.Second))

  r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("{\"status\": \"online\"}"))
  })

  r.Get("/{id}", getCulture)

  fmt.Println("Culture Generator API is online.")
	log.Fatal(http.ListenAndServe(":9913", r))
}