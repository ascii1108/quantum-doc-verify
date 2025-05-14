// api/index.go
package handler

import (
    "encoding/json"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    resp := map[string]string{
        "status": "ok",
        "message": "Quantum-Doc-Verify API is running",
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}