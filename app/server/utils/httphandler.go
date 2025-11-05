package web

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

type BaseHttpRequestPagination struct {
	Page        int           `json:"page"`
	ItemPerPage int           `json:"itemsPerPage"`
	SortBy      []interface{} `json:"sortBy"`
	GroupBy     []interface{} `json:"groupBy"`
}

type BaseHttpResponse struct {
	Result  any    `json:"result"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type Semaphore struct {
	semaCh chan struct{}
}

type BaseHttpHandler[T any] func(http.ResponseWriter, *http.Request, *T) (any, error) /* {
	//HandleFn func(http.ResponseWriter, *http.Request, *T) (any, error)
}*/

func HandleResponseErr(w http.ResponseWriter, err error) bool {
	if err != nil {
		response := BaseHttpResponse{
			Success: err == nil,
			Result:  nil,
			Error:   err.Error(),
		}
		jsonStr, _ := json.Marshal(response)
		w.Write(jsonStr)
		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Err(err)

		return true
	}

	return false
}

func (handleFn BaseHttpHandler[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, _ := io.ReadAll(r.Body)
	var requestBody T
	bodyErr := json.Unmarshal(body, &requestBody)
	if HandleResponseErr(w, bodyErr) {
		return
	}

	result, err := handleFn(w, r, &requestBody)
	response := BaseHttpResponse{
		Success: err == nil,
		Result:  result,
	}
	if err != nil {
		response.Error = err.Error()
		log.Error().Err(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	jsonStr, _ := json.Marshal(response)
	w.Write(jsonStr)
}

type JsonStringValue struct {
	Value string `json:"value"`
}

type JsonStringsValues struct {
	Values []string `json:"values"`
}
