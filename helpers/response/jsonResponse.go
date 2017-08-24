package response

import (
	"encoding/json"
	"net/http"
	"bytes"
)

func JsonResponse(data interface{}, w http.ResponseWriter) {

	json, err :=  JSONMarshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.Write(json)
}

func EmptyJsonArrayResponse(w http.ResponseWriter) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.Write([]byte("[]"))
}

func JSONMarshal(v interface{}) ([]byte, error) {
    b, err := json.Marshal(v)

    b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
    b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
    b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
    
    return b, err
}