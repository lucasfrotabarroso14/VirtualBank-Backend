package responses

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Result     interface{} `json:"result"`
}
type JWTResponse struct {
	StatusCode int         `json:"status_code"`
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Token      interface{} `json:"token"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := Response{
		StatusCode: statusCode,
		Status:     true,
		Message:    "success",
		Result:     data,
	}

	if erro := json.NewEncoder(w).Encode(response); erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

}

func Erro(w http.ResponseWriter, statusCode int, erro error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := Response{
		StatusCode: statusCode,
		Status:     false,
		Message:    erro.Error(),
		Result:     nil,
	}
	if erro := json.NewEncoder(w).Encode(response); erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}
}

func JWTtoJSON(w http.ResponseWriter, statusCode int, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := JWTResponse{
		StatusCode: statusCode,
		Status:     true,
		Message:    "success",
		Token:      data,
	}

	if erro := json.NewEncoder(w).Encode(response); erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

}
