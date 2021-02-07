# sandbox

## concurrency package

```A
pkg: sandbox/concurrency
BenchmarkCase1-3   	    2056	    519714 ns/op	     827 B/op	      11 allocs/op
BenchmarkCase2-3   	       3	 405691233 ns/op	    1072 B/op	      13 allocs/op
BenchmarkCase3-3   	       1	1013625000 ns/op	     176 B/op	       2 allocs/op
```

## gin server with Jwt

### strat 
`go run cum/server/main.go`

`curl -H "Authorization: Bearer <token>" http://localhost:8080/user`


## generate rpa pem key

`go run cmd/rsa/main.go`