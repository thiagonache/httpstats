package httpstats_test

import (
	"net/http"
	"testing"

	"github.com/thiagonache/httpstats"
)

func TestSetHTTPTrace_WithAnyMethodTracksDNSTimeOnDo(t *testing.T) {
	t.Parallel()
	s := httpstats.New()
	req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = s.SetHTTPTrace(req)
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	dns := s.DNS[0]
	if dns <= 0 {
		t.Fatalf("want DNS time to be bigger than zero, got %v", dns)
	}
}

func TestSetHTTPTrace_WithAnyMethodTracksConnectTimeOnDo(t *testing.T) {
	t.Parallel()
	s := httpstats.New()
	req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = s.SetHTTPTrace(req)
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	connect := s.Connect[0]
	if connect <= 0 {
		t.Fatalf("want Connect time to be bigger than zero, got %v", connect)
	}
}

func TestSetHTTPTrace_WithAnyMethodTracksTLSTimeOnDo(t *testing.T) {
	t.Parallel()
	s := httpstats.New()
	req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = s.SetHTTPTrace(req)
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	tls := s.TLS[0]
	if tls <= 0 {
		t.Fatalf("want TLS time to be bigger than zero, got %v", tls)
	}
}

func TestSetHTTPTrace_WithAnyMethodTracksSendTimeOnDo(t *testing.T) {
	t.Parallel()
	s := httpstats.New()
	req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = s.SetHTTPTrace(req)
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	send := s.Send[0]
	if send <= 0 {
		t.Fatalf("want send time to be bigger than zero, got %v", send)
	}
}

func TestSetHTTPTrace_WithAnyMethodTracksWaitTimeOnDo(t *testing.T) {
	t.Parallel()
	s := httpstats.New()
	req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = s.SetHTTPTrace(req)
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	wait := s.Wait[0]
	if wait <= 0 {
		t.Fatalf("want wait time to be bigger than zero, got %v", wait)
	}
}

func TestNewRequest_WithAnyMethodTracksDNSTimeOnDo(t *testing.T) {
	t.Parallel()
	s := httpstats.New()
	req, err := s.NewRequest(http.MethodDelete, "https://httpbin.org/delete", nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if s.DNS[0] <= 0 {
		t.Fatalf("want dns time to be bigger than zero, got %v", s.DNS[0])
	}
}
