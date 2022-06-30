package httpstats

import (
	"net/http"
	"net/http/httptrace"
	"time"
)

type Stats struct {
	DNS  time.Duration
	next http.RoundTripper
}

func NewHTTPStats() *Stats {
	return &Stats{next: http.DefaultTransport}
}

func (s *Stats) RoundTrip(r *http.Request) (*http.Response, error) {
	var dns time.Time
	ct := &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			dns = time.Now()
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			s.DNS = time.Since(dns)
		},
	}
	ctCtx := httptrace.WithClientTrace(r.Context(), ct)
	r = r.WithContext(ctCtx)
	return s.next.RoundTrip(r)
}

func (s *Stats) Get(url string) (*http.Response, error) {
	c := &http.Client{Transport: s}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return &http.Response{}, err
	}
	res, err := c.Do(req)
	if err != nil {
		return &http.Response{}, err
	}
	return res, nil
}
