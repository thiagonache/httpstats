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
		t.Errorf("want DNS time to be bigger than 0, got %v", s.DNS[0])
	}
}

func TestNewRequestWithMethodGetTracksConnectTime(t *testing.T) {
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
	if s.Connect[0] <= 0 {
		t.Errorf("want Connect time to be bigger than 0, got %v", s.Connect[0])
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
		t.Errorf("want DNS time to be bigger than 0, got %v", s.DNS[0])
	}
}

func TestGetTracksTLSTime(t *testing.T) {
	t.Parallel()
	s := httpstats.NewHTTPStats()
	_, err := s.Get("https://httpbin.org/get")
	if err != nil {
		t.Fatal(err)
	}
	if s.TLS[0] <= 0 {
		t.Errorf("want TLS time to be bigger than 0, got %v", s.TLS[0])
	}
}

func TestGetTracksSendTime(t *testing.T) {
	t.Parallel()
	s := httpstats.NewHTTPStats()
	_, err := s.Get("https://httpbin.org/get")
	if err != nil {
		t.Fatal(err)
	}
	if s.Send[0] <= 0 {
		t.Errorf("want Send time to be bigger than 0, got %v", s.Send[0])
	}
}

func TestGetTracksWaitTime(t *testing.T) {
	t.Parallel()
	s := httpstats.NewHTTPStats()
	_, err := s.Get("https://httpbin.org/get")
	if err != nil {
		t.Fatal(err)
	}
	if s.Wait[0] <= 0 {
		t.Errorf("want Wait time to be bigger than 0, got %v", s.Wait[0])
	}
}
