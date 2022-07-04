package httpstats

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"time"
)

type Stats struct {
	client   *http.Client
	Connect  []time.Duration
	DNS      []time.Duration
	next     http.RoundTripper
	Send     []time.Duration
	TLS      []time.Duration
	Total    []time.Duration
	Transfer []time.Duration
	Wait     []time.Duration
}

func NewHTTPStats(opts ...Option) *Stats {
	stats := &Stats{
		client: &http.Client{},
		next:   http.DefaultTransport,
	}
	for _, o := range opts {
		o(stats)
	}
	if stats.client.Transport != nil {
		stats.next = stats.client.Transport
	}
	return stats
}

func (s *Stats) RoundTrip(r *http.Request) (*http.Response, error) {
	var (
		connectStart time.Time
		dnsStart     time.Time
		reqStart     time.Time
		sendStart    time.Time
		tlsStart     time.Time
		waitStart    time.Time
	)
	ct := &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			reqStart = time.Now()
		},
		DNSStart: func(info httptrace.DNSStartInfo) {
			dnsStart = time.Now()
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			s.DNS = append(s.DNS, time.Since(dnsStart))
		},
		ConnectStart: func(network, addr string) {
			connectStart = time.Now()
		},
		ConnectDone: func(network, addr string, err error) {
			if err == nil {
				s.Connect = append(s.Connect, time.Since(connectStart))
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
			sendStart = time.Now()
		},
		WroteHeaders: func() {
			s.Send = append(s.Send, time.Since(sendStart))
		},
		WroteRequest: func(wri httptrace.WroteRequestInfo) {
			waitStart = time.Now()
		},
		GotFirstResponseByte: func() {
			s.Wait = append(s.Wait, time.Since(waitStart))
		},
	}
	ctCtx := httptrace.WithClientTrace(r.Context(), ct)
	r = r.WithContext(ctCtx)
	res, err := s.next.RoundTrip(r)
	transferStart := time.Now()
	body, errBody := ioutil.ReadAll(res.Body)
	if errBody == nil {
		res.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}
	s.Transfer = append(s.Transfer, time.Since(transferStart))
	s.Total = append(s.Total, time.Since(reqStart))
	return res, err
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

type Option func(*Stats)

// WithHTTPClient is the functional option to set a custom http.Client while
// initializing a new Stats object
func WithHTTPClient(client *http.Client) Option {
	return func(s *Stats) {
		s.client = client
	}
}
