# httpstats

```go
import "github.com/thiagonache/httpstats"
```

The idea of this package is to provide an easy way to collect detailed HTTP
metrics from the client. It exports a function that injects hooks using
[HTTPTrace](https://pkg.go.dev/net/http/httptrace) Package. If you want a better 
understand about how HTTPTrace works read [my blog post](https://hi.thiagonbcarvalho.com/en/posts/httptrace/).

You just need to instantiate the object and call the function `SetHTTPTrace`
after setting up the `http.Request`. The trace will be re-used for all requests
done using that request object.

## Example

Usual http call code

```go
c := &http.Client{}
req, err := http.NewRequest(http.MethodGet, "https://httpbin.org/get", nil)
if err != nil {
    log.Fatal(err)
}
res, err := c.Do(req)
if err != nil {
    log.Fatal(err)
}
body, err := io.ReadAll(res.Body)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(body))
```

Output

```json
{
  "args": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/2.0",
    "X-Amzn-Trace-Id": "Root=1-62bdfa65-252f4f6244194d1257217609"
  },
  "origin": "201.95.218.50",
  "url": "https://httpbin.org/get"
}
```

Using this package to get detailed metrics.

```go
s := httpstats.New()
c := &http.Client{}
req, err := s.NewRequest(http.MethodGet, "https://httpbin.org/get", nil)
...
fmt.Println("DNS:", s.DNS[0])
fmt.Println("Connect:", s.Connect[0])
```

Output

```json
{
  "args": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/2.0",
    "X-Amzn-Trace-Id": "Root=1-62bdfaef-6a1f4a8533030f0f78dbb4e6"
  },
  "origin": "201.95.218.50",
  "url": "https://httpbin.org/get"
}

DNS: 9.454983ms
Connect: 147.036392ms
```
