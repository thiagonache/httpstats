package httpstats

import (
	"crypto/tls"
	"net/http"
	"net/http/httptrace"
	"sync"
	"time"
)

type Stats struct {
	mu           *sync.Mutex
	Connect      []time.Duration
	ConnectStart time.Time
	DNS          []time.Duration
	DNSStart     time.Time
	Send         []time.Duration
	SendStart    time.Time
	TLS          []time.Duration
	TLSStart     time.Time
	Wait         []time.Duration
	WaitStart    time.Time
}

func New() *Stats {
	return &Stats{
		mu: &sync.Mutex{},
	}
}

func (s *Stats) dnsStart(di httptrace.DNSStartInfo) {
	s.mu.Lock()
	s.DNSStart = time.Now()
	s.mu.Unlock()
}

func (s *Stats) dnsDone(di httptrace.DNSDoneInfo) {
	s.mu.Lock()
	s.DNS = append(s.DNS, time.Since(s.DNSStart))
	s.mu.Unlock()
}

func (s *Stats) connectStart(netProto, addr string) {
	s.mu.Lock()
	s.ConnectStart = time.Now()
	s.mu.Unlock()
}

func (s *Stats) connectDone(netProto, addr string, err error) {
	if err == nil {
		s.mu.Lock()
		s.Connect = append(s.Connect, time.Since(s.ConnectStart))
		s.mu.Unlock()
	}
}

func (s *Stats) tlsStart() {
	s.mu.Lock()
	s.TLSStart = time.Now()
	s.mu.Unlock()
}

func (s *Stats) tlsDone(cs tls.ConnectionState, err error) {
	if err == nil {
		s.mu.Lock()
		s.TLS = append(s.TLS, time.Since(s.TLSStart))
		s.mu.Unlock()
	}
}

func (s *Stats) wroteHeaderField(key string, value []string) {
	s.mu.Lock()
	if s.SendStart.IsZero() {
		s.SendStart = time.Now()
	}
	s.mu.Unlock()
}

func (s *Stats) wroteHeaders() {
	s.mu.Lock()
	s.Send = append(s.Send, time.Since(s.SendStart))
	s.mu.Unlock()
}

func (s *Stats) wroteRequest(wri httptrace.WroteRequestInfo) {
	s.mu.Lock()
	s.WaitStart = time.Now()
	s.mu.Unlock()
}

func (s *Stats) gotFirstResponseByte() {
	s.mu.Lock()
	s.Wait = append(s.Wait, time.Since(s.WaitStart))
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
