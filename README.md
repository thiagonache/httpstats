# httpstats

```go
import "github.com/thiagonache/httpstats"

func main() {
  s := httpstats.NewHTTPStats()
  c := &http.Client{Transport: s}
  ...
}
```

The idea of this package is to provide an easy way to collect detailed HTTP
metrics from the client. It implements a HTTP Round Tripper that uses HTTPTrace Package.

You just need to instantiate the object and inject it in the Transport field of
your HTTP Client. In case you already implement a custom HTTP Round Tripper, this
package will provide a functional option to set the next round tripper instead
of calling the default one.

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
s := httpstats.NewHTTPStats()
c := &http.Client{Transport: s}
...
fmt.Println(s.DNS)
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

[9.454983ms]
```
