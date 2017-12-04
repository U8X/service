package main

import (
	"encoding/json"
	"net/http"

	u8xhttp "github.com/U8X/service/http"
	"github.com/U8X/service/shorten"
)

type server struct {
	Prefix  string
	Service shorten.Interface
}

func (s *server) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/v1/shorten" {
		s.doShorten(w, request)
		return
	}
	s.doRedirect(w, request)
}

func (s *server) doShorten(w http.ResponseWriter, request *http.Request) {
	l := request.URL.Query().Get("long_url")
	if len(l) == 0 {
		http.Error(w, "缺少long_url", http.StatusBadRequest)
		return
	}
	v := s.Service.Create(l)
	resp := u8xhttp.Response{Data: s.Prefix + v}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resp)
}

func (s *server) doRedirect(w http.ResponseWriter, request *http.Request) {
	short := request.URL.Path[1:]
	long, err := s.Service.Expand(short)
	if err != nil {
		http.NotFound(w, request)
		return
	}
	w.Header().Set("Location", long)
	w.WriteHeader(http.StatusMovedPermanently)
	w.Write([]byte("Redirecting..."))
}
