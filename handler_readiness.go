package main

import (
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	type status struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, 200, status{
		Status: "success",
	})
}