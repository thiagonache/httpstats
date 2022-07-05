package httpstats

import (
	"crypto/tls"
	"net/http"
	"net/http/httptrace"
	"sync"
	"time"
)

type Stats struct {
	mu       *sync.Mutex
	Connect  []time.Duration
	ConnectS time.Time
	DNS      []time.Duration
	DNSS     time.Time
	Send     []time.Duration
	SendS    time.Time
	TLS      []time.Duration
	TLSS     time.Time
	Wait     []time.Duration
	WaitS    time.Time
}

func NewHTTPStats() *Stats {
	return &Stats{mu: &sync.Mutex{}}
}

func (s *Stats) SetHTTPTrace(r *http.Request) *http.Request {
	ct := &httptrace.ClientTrace{
		DNSStart: func(di httptrace.DNSStartInfo) {
			s.mu.Lock()
			s.DNSS = time.Now()
			s.mu.Unlock()
		},
		DNSDone: func(di httptrace.DNSDoneInfo) {
			s.mu.Lock()
			s.DNS = append(s.DNS, time.Since(s.DNSS))
			s.mu.Unlock()
		},
		ConnectStart: func(network, addr string) {
			s.mu.Lock()
			s.ConnectS = time.Now()
			s.mu.Unlock()
		},
		ConnectDone: func(network, addr string, err error) {
			if err == nil {
				s.mu.Lock()
				s.Connect = append(s.Connect, time.Since(s.ConnectS))
				s.mu.Unlock()
			}
		},
		TLSHandshakeStart: func() {
			s.mu.Lock()
			s.TLSS = time.Now()
			s.mu.Unlock()
		},
		TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
			if err == nil {
				s.mu.Lock()
				s.TLS = append(s.TLS, time.Since(s.TLSS))
				s.mu.Unlock()
			}
		},
		WroteHeaderField: func(key string, value []string) {
			s.mu.Lock()
			s.SendS = time.Now()
			s.mu.Unlock()
		},
		WroteHeaders: func() {
			s.mu.Lock()
			s.Send = append(s.Send, time.Since(s.SendS))
			s.mu.Unlock()
		},
		WroteRequest: func(wri httptrace.WroteRequestInfo) {
			s.mu.Lock()
			s.WaitS = time.Now()
			s.mu.Unlock()
		},
		GotFirstResponseByte: func() {
			s.mu.Lock()
			s.Wait = append(s.Wait, time.Since(s.WaitS))
			s.mu.Unlock()
		},
	}
	ctCtx := httptrace.WithClientTrace(r.Context(), ct)
	r = r.WithContext(ctCtx)
	return r
}
