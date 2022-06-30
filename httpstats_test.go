package httpstats_test

import (
	"net/http"
	"testing"

	"github.com/thiagonache/httpstats"
)

func TestNewRequestWithMethodGetTracksDNSTime(t *testing.T) {
	t.Parallel()
	s := httpstats.NewHTTPStats()
	c := &http.Client{Transport: s}
	req, err := http.NewRequest(http.MethodGet, "https://httpbin.org/get", nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if s.DNS[0] <= 0 {
		t.Errorf("want DNS time to be bigger than 0, got %v", s.DNS)
	}
}

func TestGetTracksDNSTime(t *testing.T) {
	t.Parallel()
	s := httpstats.NewHTTPStats()
	_, err := s.Get("https://httpbin.org/get")
	if err != nil {
		t.Fatal(err)
	}
	if s.DNS[0] <= 0 {
		t.Errorf("want DNS time to be bigger than 0, got %v", s.DNS)
	}
}
