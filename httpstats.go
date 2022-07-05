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

func (s *Stats) dnsStart(di httptrace.DNSStartInfo) {
	s.mu.Lock()
	s.DNSS = time.Now()
	s.mu.Unlock()
}

func (s *Stats) dnsDone(di httptrace.DNSDoneInfo) {
	s.mu.Lock()
	s.DNS = append(s.DNS, time.Since(s.DNSS))
	s.mu.Unlock()
}

func (s *Stats) connectStart(netProto, addr string) {
	s.mu.Lock()
	s.ConnectS = time.Now()
	s.mu.Unlock()
}

func (s *Stats) connectDone(netProto, addr string, err error) {
	if err == nil {
		s.mu.Lock()
		s.Connect = append(s.Connect, time.Since(s.ConnectS))
		s.mu.Unlock()
	}
}

func (s *Stats) tlsStart() {
	s.mu.Lock()
	s.TLSS = time.Now()
	s.mu.Unlock()
}

func (s *Stats) tlsDone(cs tls.ConnectionState, err error) {
	if err == nil {
		s.mu.Lock()
		s.TLS = append(s.TLS, time.Since(s.TLSS))
		s.mu.Unlock()
	}
}

func (s *Stats) wroteHeaderField(key string, value []string) {
	s.mu.Lock()
	if s.SendS.IsZero() {
		s.SendS = time.Now()
	}
	s.mu.Unlock()
}

func (s *Stats) wroteHeaders() {
	s.mu.Lock()
	s.Send = append(s.Send, time.Since(s.SendS))
	s.mu.Unlock()
}

func (s *Stats) wroteRequest(wri httptrace.WroteRequestInfo) {
	s.mu.Lock()
	s.WaitS = time.Now()
	s.mu.Unlock()
}

func (s *Stats) gotFirstResponseByte() {
	s.mu.Lock()
	s.Wait = append(s.Wait, time.Since(s.WaitS))
	s.mu.Unlock()
}

func (s *Stats) SetHTTPTrace(r *http.Request) *http.Request {
	ct := &httptrace.ClientTrace{
		DNSStart:             s.dnsStart,
		DNSDone:              s.dnsDone,
		ConnectStart:         s.connectStart,
		ConnectDone:          s.connectDone,
		TLSHandshakeStart:    s.tlsStart,
		TLSHandshakeDone:     s.tlsDone,
		WroteHeaderField:     s.wroteHeaderField,
		WroteHeaders:         s.wroteHeaders,
		WroteRequest:         s.wroteRequest,
		GotFirstResponseByte: s.gotFirstResponseByte,
	}
	ctCtx := httptrace.WithClientTrace(r.Context(), ct)
	r = r.WithContext(ctCtx)
	return r
}
