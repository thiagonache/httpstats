package httpstats_test

import (
	"net/http"
	"testing"

	"github.com/thiagonache/httpstats"
)

func TestSetHTTPTrace_TracksDNSTime(t *testing.T) {
	t.Parallel()
	hstats := httpstats.NewHTTPStats()
	req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = hstats.SetHTTPTrace(req)
	c := &http.Client{}
	_, err = c.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	dns := hstats.DNS[0]
	if dns <= 0 {
		t.Fatalf("want DNS time to be bigger than zero, got %v", dns)
	}
}

func TestSetHTTPTrace_TracksConnectTime(t *testing.T) {
	t.Parallel()
	hstats := httpstats.NewHTTPStats()
	req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = hstats.SetHTTPTrace(req)
	c := &http.Client{}
	_, err = c.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	connect := hstats.Connect[0]
	if connect <= 0 {
		t.Fatalf("want Connect time to be bigger than zero, got %v", connect)
	}
}

func TestSetHTTPTrace_TracksTLSTime(t *testing.T) {
	t.Parallel()
	hstats := httpstats.NewHTTPStats()
	req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = hstats.SetHTTPTrace(req)
	c := &http.Client{}
	_, err = c.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	tls := hstats.TLS[0]
	if tls <= 0 {
		t.Fatalf("want TLS time to be bigger than zero, got %v", tls)
	}
}

func TestSetHTTPTrace_TracksSendTime(t *testing.T) {
	t.Parallel()
	hstats := httpstats.NewHTTPStats()
	req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = hstats.SetHTTPTrace(req)
	c := &http.Client{}
	_, err = c.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	send := hstats.Send[0]
	if send <= 0 {
		t.Fatalf("want send time to be bigger than zero, got %v", send)
	}
}

func TestSetHTTPTrace_TracksWaitTime(t *testing.T) {
	t.Parallel()
	hstats := httpstats.NewHTTPStats()
	req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = hstats.SetHTTPTrace(req)
	c := &http.Client{}
	_, err = c.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	wait := hstats.Wait[0]
	if wait <= 0 {
		t.Fatalf("want wait time to be bigger than zero, got %v", wait)
	}
}
