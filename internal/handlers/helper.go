package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/anuragprafulla/bullet/pkg/errors"
)

type Response interface {
	Json() []byte
	StatusCode() int
}

func WriteResponse(w http.ResponseWriter, res Response) {
	w.WriteHeader(res.StatusCode())
	_, _ = w.Write(res.Json())
}

func WriteError(w http.ResponseWriter, err error) {
	res, ok := err.(*errors.Error)
	if !ok {
		log.Println(err)
		res = errors.ErrorInternalServer
	}
	WriteResponse(w, res)
}

func IntFromString(w http.ResponseWriter, v string) (int, error) {
	if v == "" {
		return 0, nil
	}
	res, err := strconv.Atoi(v)
	if err != nil {
		log.Println(err)
		WriteError(w, errors.ErrorBadRequest)
	}
	return res, err
}

func Unmarshal(w http.ResponseWriter, data []byte, v interface{}) error {
	if d := string(data); d == "null" || d == "" {
		WriteError(w, errors.ErrorBadRequest)
		return errors.ErrorBadRequest
	}
	err := json.Unmarshal(data, v)
	if err != nil {
		log.Println(err)
		WriteError(w, errors.ErrorBadRequest)
	}
	return err
}
