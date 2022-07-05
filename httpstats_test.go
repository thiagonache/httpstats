package httpstats_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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
	if len(s.DNS) == 0 {
		t.Fatal("no DNS time recorded")
	}
	if s.DNS[0] <= 0 {
		t.Fatalf("want DNS time to be bigger than zero, got %v", s.DNS[0])
	}
}

func TestSetHTTPTrace_WithAnyMethodTracksConnectTimeOnDo(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "OK")
	}))
	s := httpstats.New()
	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	req = s.SetHTTPTrace(req)
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if len(s.Connect) == 0 {
		t.Fatal("no connect time recorded")
	}
	if s.Connect[0] <= 0 {
		t.Fatalf("want connect time to be bigger than zero, got %v", s.Connect[0])
	}
}

func TestSetHTTPTrace_WithAnyMethodTracksTLSTimeOnDo(t *testing.T) {
	t.Parallel()
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "OK")
	}))
	s := httpstats.New()
	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	req = s.SetHTTPTrace(req)
	_, err = server.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if len(s.TLS) == 0 {
		t.Fatal("no TLS time recorded")
	}
	if s.TLS[0] <= 0 {
		t.Fatalf("want TLS time to be bigger than zero, got %v", s.TLS[0])
	}
}

func TestSetHTTPTrace_WithAnyMethodTracksSendTimeOnDo(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "OK")
	}))
	s := httpstats.New()
	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	req = s.SetHTTPTrace(req)
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if len(s.Send) == 0 {
		t.Fatal("no send time recorded")
	}
	if s.Send[0] <= 0 {
		t.Fatalf("want send time to be bigger than zero, got %v", s.Send[0])
	}
}

func TestSetHTTPTrace_WithAnyMethodTracksWaitTimeOnDo(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "OK")
	}))
	s := httpstats.New()
	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	req = s.SetHTTPTrace(req)
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if len(s.Wait) == 0 {
		t.Fatal("no wait time recorded")
	}
	if s.Wait[0] <= 0 {
		t.Fatalf("want wait time to be bigger than zero, got %v", s.Wait[0])
	}
}

func TestNewRequest_WithAnyMethodTracksConnectTimeOnDo(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "OK")
	}))
	s := httpstats.New()
	req, err := s.NewRequest(http.MethodDelete, server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if len(s.Connect) == 0 {
		t.Fatal("no connect time recorded")
	}
	if s.Connect[0] <= 0 {
		t.Fatalf("want connect time to be bigger than zero, got %v", s.Connect[0])
	}
}

func TestTwoRequestsTracksMinTwoConnectTime(t *testing.T) {
	t.Parallel()
	s1 := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "OK")
	}))
	s2 := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "OK")
	}))
	s := httpstats.New()
	req, err := s.NewRequest(http.MethodGet, s1.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	req, err = s.NewRequest(http.MethodGet, s2.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if len(s.Connect) == 0 {
		t.Fatal("no connect time recorded")
	}
	if len(s.Connect) < 2 {
		t.Fatalf("want min of two connect time recorded but got %d: %q", len(s.Connect), s.Connect)
	}
}
