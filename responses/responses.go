package responses

import (
	"encoding/json"
	"net/http"
)

// resposta que o result é uma interface
type Response struct {
	StatusCode int         `json:"status_code"`
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Result     interface{} `json:"result"`
}

// respostas quando for retornar o token
type JWTResponse struct {
	StatusCode int         `json:"status_code"`
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Token      interface{} `json:"token"`
}

type JSONOwnResponse struct {
	StatusCode int    `json:"status_code"`
	Status     bool   `json:"status"`
	Message    string `json:"message"`
	Result     string `json:"result"`
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

func JWTjsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {

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

func MakeJSONResponse(w http.ResponseWriter, statusCode int, data string) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := JSONOwnResponse{
		StatusCode: statusCode,
		Status:     true,
		Message:    "success",
		Result:     data,
	}

	if erro := json.NewEncoder(w).Encode(response); erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
	}

}
