package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
)

type Result struct {
	Greetings string `json:"greetings,omitempty"`
	Name      string `json:"name,omitempty"`
}

type Answer struct {
	Status string `json:"status"`
	Result Result `json:"result"`
}

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name != "" && !regexp.MustCompile("^[a-zA-Z]+$").MatchString(name) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(Answer{
				Status: "error",
				Result: Result{},
			})
			return
		}
		next(w, r)
	}
}

func RPC(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "stranger"
		}
		ctx := context.WithValue(r.Context(), "name", name)

		r = r.WithContext(ctx)
		next(w, r)
	}
}

func StrangerHandler(w http.ResponseWriter, r *http.Request) {

	name := r.Context().Value("name").(string)
	output := Result{
		Greetings: "hello",
		Name:      name,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Answer{
		Status: "ok",
		Result: output,
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	finalHandler := SetDefaultName(RPC(Sanitize(StrangerHandler)))
	finalHandler(w, r)
}

func main() {
	http.HandleFunc("/", HelloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
