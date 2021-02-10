package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type server struct {
	URL string
	reverseProxy *httputil.ReverseProxy
	health bool
}

func newServer(urlString string) *server {
	u, _ := url.Parse(urlString)
	rproxy := httputil.NewSingleHostReverseProxy(u)

	return &server{
		URL : urlString,
		reverseProxy : rproxy,
		health : true,
	}
}

func (s *server) checkHealth() bool {
	resp, err := http.Head(s.URL)
	if err != nil {
		s.health = false
		return s.health
	}

	if resp.StatusCode != http.StatusOK {
		s.health = false
		return s.health
	}
	s.health = true
	return s.health
}

