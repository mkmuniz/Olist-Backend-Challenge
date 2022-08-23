package httperror

import (
	"encoding/json"
	"net/http"
)

var resp map[string]any

func RequestError(w http.ResponseWriter, r *http.Request, res interface{}, err error) {

	if err != nil {
		resp = map[string]any{
			"status":  500,
			"message": "Error on request",
			"data":    res,
		}
	} else {
		resp = map[string]any{
			"status":  200,
			"message": "Sucess",
			"data":    res,
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func DecodeError(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		resp = map[string]any{
			"status":  500,
			"message": "Error on decode body",
		}

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
