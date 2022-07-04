package httpstats

import (
	"net/http"
	"net/http/httptrace"
	"time"
)

type Stats struct {
	DNS []time.Duration
}

func NewHTTPStats() *Stats {
	return &Stats{}
}

func (s *Stats) SetHTTPTrace(r *http.Request) *http.Request {
	var dns time.Time
	ct := &httptrace.ClientTrace{
		DNSStart: func(di httptrace.DNSStartInfo) {
			dns = time.Now()
		},
		DNSDone: func(di httptrace.DNSDoneInfo) {
			s.DNS = append(s.DNS, time.Since(dns))
		},
	}
	ctCtx := httptrace.WithClientTrace(r.Context(), ct)
	r = r.WithContext(ctCtx)
	return r
}

// func (s *Stats) RoundTrip(r *http.Request) (*http.Response, error) {
// 	var (
// 		connectStart time.Time
// 		dnsStart     time.Time
// 		reqStart     time.Time
// 		sendStart    time.Time
// 		tlsStart     time.Time
// 		waitStart    time.Time
// 	)
// 	ct := &httptrace.ClientTrace{
// 		GetConn: func(hostPort string) {
// 			reqStart = time.Now()
// 		},
// 		DNSStart: func(info httptrace.DNSStartInfo) {
// 			dnsStart = time.Now()
// 		},
// 		DNSDone: func(info httptrace.DNSDoneInfo) {
// 			s.DNS = append(s.DNS, time.Since(dnsStart))
// 		},
// 		ConnectStart: func(network, addr string) {
// 			connectStart = time.Now()
// 		},
// 		ConnectDone: func(network, addr string, err error) {
// 			if err == nil {
// 				s.Connect = append(s.Connect, time.Since(connectStart))
// 			}
// 		},
// 		TLSHandshakeStart: func() {
// 			tlsStart = time.Now()
// 		},
// 		TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
// 			if err == nil {
// 				s.TLS = append(s.TLS, time.Since(tlsStart))
// 			}
// 		},
// 		WroteHeaderField: func(key string, value []string) {
// 			sendStart = time.Now()
// 		},
// 		WroteHeaders: func() {
// 			s.Send = append(s.Send, time.Since(sendStart))
// 		},
// 		WroteRequest: func(wri httptrace.WroteRequestInfo) {
// 			waitStart = time.Now()
// 		},
// 		GotFirstResponseByte: func() {
// 			s.Wait = append(s.Wait, time.Since(waitStart))
// 		},
// 	}
// 	ctCtx := httptrace.WithClientTrace(r.Context(), ct)
// 	r = r.WithContext(ctCtx)
// 	res, err := s.next.RoundTrip(r)
// 	transferStart := time.Now()
// 	body, errBody := ioutil.ReadAll(res.Body)
// 	if errBody == nil {
// 		res.Body = ioutil.NopCloser(bytes.NewBuffer(body))
// 	}
// 	s.Transfer = append(s.Transfer, time.Since(transferStart))
// 	s.Total = append(s.Total, time.Since(reqStart))
// 	return res, err
// }

// func (s *Stats) Get(url string) (*http.Response, error) {
// 	c := &http.Client{Transport: s}
// 	req, err := http.NewRequest(http.MethodGet, url, nil)
// 	if err != nil {
// 		return &http.Response{}, err
// 	}
// 	res, err := c.Do(req)
// 	if err != nil {
// 		return &http.Response{}, err
// 	}
// 	return res, nil
// }

// type Option func(*Stats)

// // WithHTTPClient is the functional option to set a custom http.Client while
// // initializing a new Stats object
// func WithHTTPClient(client *http.Client) Option {
// 	return func(s *Stats) {
// 		s.client = client
// 	}
// }
