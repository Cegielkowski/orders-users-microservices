package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

type heartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

// HeartbeatController Verify if it's all set.
func HeartbeatController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
	if err != nil {
		log.Print(err.Error())
	}
}
