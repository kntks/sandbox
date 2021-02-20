# sandbox

## Usage
1, vscodeで`sandbox/golang`ディレクトリをルートとしてプロジェクトをオープンする
2, `command + shift + p`でコマンドパレットを開く
3, `ReOpen in Container`で開く

## concurrency package

```
pkg: sandbox/concurrency
BenchmarkCase1-3   	    2056	    519714 ns/op	     827 B/op	      11 allocs/op
BenchmarkCase2-3   	       3	 405691233 ns/op	    1072 B/op	      13 allocs/op
BenchmarkCase3-3   	       1	1013625000 ns/op	     176 B/op	       2 allocs/op
```

## client package
```bash
goos: linux
goarch: amd64
pkg: sandbox/client
BenchmarkCase11-3   	     279	   4252352 ns/op	  135790 B/op	    1002 allocs/op
BenchmarkCase12-3   	     244	   4540078 ns/op	  136554 B/op	    1013 allocs/op
```

## gin server with Jwt

### start
`go run cum/server/main.go`

`curl -H "Authorization: Bearer <token>" http://localhost:8080/user`


## generate rpa pem key

`go run cmd/rsa/main.go`