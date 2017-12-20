package http

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/U8X/service/shorten"
)

// Server U8X service http server
type Server struct {
	Prefix  string
	Service shorten.Interface
}

func (s *Server) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/v1/shorten" {
		s.doShorten(w, request)
		return
	}
	if request.URL.Path == "/v1/expand" {
		s.doExpand(w, request)
		return
	}
	s.doRedirect(w, request)
}

func (s *Server) doShorten(w http.ResponseWriter, request *http.Request) {
	l := request.URL.Query().Get("long_url")
	if len(l) == 0 {
		s.response(w, CodeCommonError, "", "缺少long_url")
		return
	}
	v := s.Service.Create(l)
	s.response(w, 0, s.Prefix+v, "")
}

func (s *Server) doExpand(w http.ResponseWriter, request *http.Request) {
	shortUrl, err := url.Parse(request.URL.Query().Get("short_url"))
	if err != nil {
		s.response(w, CodeCommonError, "", err.Error())
		return
	}
	short := shortUrl.Path[1:]
	long, err := s.Service.Expand(short)
	if err != nil {
		s.response(w, CodeCommonError, "", err.Error())
		return
	}
	s.response(w, 0, long, "")
}

func (s *Server) doRedirect(w http.ResponseWriter, request *http.Request) {
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

func (s *Server) response(w http.ResponseWriter, code int, data interface{}, errMsg string) {
	w.WriteHeader(http.StatusOK)
	resp := Response{Code: code, Data: data, ErrMsg: errMsg}
	json.NewEncoder(w).Encode(&resp)
}
