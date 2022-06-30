# httpstats

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

Getting HTTP detailed metrics

```go
s := httpstats.NewHTTPStats()
c := &http.Client{Transport: s}
...
fmt.Println(s.DNS)
```

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
