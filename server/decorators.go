package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/pabloos/http/greet"
)

func POST(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func Debug(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer h.ServeHTTP(w, r)

		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(dump))
	}
}

func Delay(delay time.Duration, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer h.ServeHTTP(w, r)

		time.Sleep(delay)
	}
}

func Cache(h http.HandlerFunc) http.HandlerFunc {
	messages := make(map[string]greet.Greet)

	return func(w http.ResponseWriter, r *http.Request) {
		// Read body into buffer to avoid consuming it - https://stackoverflow.com/questions/23070876/reading-body-of-http-request-without-modifying-request-state
		bufferBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}

		r.Body = ioutil.NopCloser(bytes.NewBuffer(bufferBody))
		defer h.ServeHTTP(w, r)

		var parsedBody greet.Greet
		body := ioutil.NopCloser(bytes.NewBuffer(bufferBody))
		json.NewDecoder(body).Decode(&parsedBody)
		cacheData, found := messages[parsedBody.Name]
		if found {
			fmt.Fprintf(w, "%s, from %s was found in cache\n", cacheData.Name, cacheData.Location)
		} else {
			messages[parsedBody.Name] = parsedBody
		}
	}
}
