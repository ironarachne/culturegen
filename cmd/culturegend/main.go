package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ironarachne/culturegen"
	"github.com/ironarachne/random"
)

func getCulture(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var newCulture culturegen.Culture

	random.SeedFromString(id)

	newCulture = culturegen.GenerateCulture()

	json.NewEncoder(w).Encode(newCulture)
}

func getCultureRandom(w http.ResponseWriter, r *http.Request) {
	var newCulture culturegen.Culture

	rand.Seed(time.Now().UnixNano())

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

	r.Get("/", getCultureRandom)
	r.Get("/{id}", getCulture)

	fmt.Println("Culture Generator API is online.")
	log.Fatal(http.ListenAndServe(":9913", r))
}
