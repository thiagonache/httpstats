package httpstats

import (
	"crypto/tls"
	"net/http"
	"net/http/httptrace"
	"time"
)

type Stats struct {
	Connect []time.Duration
	DNS     []time.Duration
	Send    []time.Duration
	TLS     []time.Duration
	Wait    []time.Duration
}

func NewHTTPStats() *Stats {
	return &Stats{}
}

func (s *Stats) SetHTTPTrace(r *http.Request) *http.Request {
	var (
		connect  time.Time
		dns      time.Time
		send     time.Time
		tlsStart time.Time
		wait     time.Time
	)
	ct := &httptrace.ClientTrace{
		DNSStart: func(di httptrace.DNSStartInfo) {
			dns = time.Now()
		},
		DNSDone: func(di httptrace.DNSDoneInfo) {
			s.DNS = append(s.DNS, time.Since(dns))
		},
		ConnectStart: func(network, addr string) {
			connect = time.Now()
		},
		ConnectDone: func(network, addr string, err error) {
			if err == nil {
				s.Connect = append(s.Connect, time.Since(connect))
			}
		},
		TLSHandshakeStart: func() {
			tlsStart = time.Now()
		},
		TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
			if err == nil {
				s.TLS = append(s.TLS, time.Since(tlsStart))
			}
		},
		WroteHeaderField: func(key string, value []string) {
			send = time.Now()
		},
		WroteHeaders: func() {
			s.Send = append(s.Send, time.Since(send))
		},
		WroteRequest: func(wri httptrace.WroteRequestInfo) {
			wait = time.Now()
		},
		GotFirstResponseByte: func() {
			s.Wait = append(s.Wait, time.Since(wait))
		},
	}
	ctCtx := httptrace.WithClientTrace(r.Context(), ct)
	r = r.WithContext(ctCtx)
	return r
}
