package controller

import (
	"encoding/json"
	"net/http"
)

func HealthCheckHandle(w http.ResponseWriter, req *http.Request) {

	var data = map[string]interface{}{
		"status":  "up",
		"code": 200,
	}
	json.NewEncoder(w).Encode(data)
}