package httpstats_test

import (
	"net/http"
	"testing"

	"github.com/thiagonache/httpstats"
)

func TestNewHTTPStats_CreatesEmptyPointerObject(t *testing.T) {
	t.Parallel()
	want := &httpstats.Stats{}
	got := httpstats.NewHTTPStats()
	if want != got {
		t.Errorf("want a pointer to an empty stats object (%v), got %v", want, got)
	}
}

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
