package httpstats

import (
	"net/http"
	"net/http/httptrace"
	"sync"
	"time"
)

type Stats struct {
	mu   *sync.Mutex
	DNS  []time.Duration
	next http.RoundTripper
}

func NewHTTPStats() *Stats {
	return &Stats{
		next: http.DefaultTransport,
		mu:   &sync.Mutex{},
	}
}

func (s *Stats) RecordDNSTime(took time.Duration) {
	s.mu.Lock()
	s.DNS = append(s.DNS, took)
	s.mu.Unlock()
}

func (s *Stats) RoundTrip(r *http.Request) (*http.Response, error) {
	var dns time.Time
	ct := &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			dns = time.Now()
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			s.RecordDNSTime(time.Since(dns))
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
