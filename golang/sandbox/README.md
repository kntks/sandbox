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

### pipeline channelとSlice

https://blog.golang.org/pipelines
```
/go/src # go run cmd/concurrent/main.go 
sliece:  1.5071818s
pipeline:  30.8µs
```

## client package
```bash
goos: linux
goarch: amd64
pkg: sandbox/client
BenchmarkCase11-3   	     279	   4252352 ns/op	  135790 B/op	    1002 allocs/op
BenchmarkCase12-3   	     244	   4540078 ns/op	  136554 B/op	    1013 allocs/op
```

##  mutex package

sync.RWMutexを `Lock()`, `UnLock()`でラップしている[RWMutex.RLocker](https://golang.org/pkg/sync/#RWMutex.RLocker)がある  
これは[Locker](https://golang.org/pkg/sync/#Locker)を実装している

```go
type Locker interface {
    Lock()
    Unlock()
}
```

コード自体は LockとUnLockをcallしているだけだがbenchmarkを取ると差が出ていることがわかる  
readするだけならRWLockを使った方が良い
```
goos: linux
goarch: amd64
pkg: sandbox/mutex
BenchmarkMyMutex-3     	  208287	      5438 ns/op	     532 B/op	       2 allocs/op
BenchmarkMyRWMutex-3   	 1000000	      1640 ns/op	      25 B/op	       0 allocs/op
PASS
```

```bash
go run cmd/mutex/main.go 
```

## gin server with Jwt

### start
`go run cum/server/main.go`

`curl -H "Authorization: Bearer <token>" http://localhost:8080/user`


## generate rpa pem key

`go run cmd/rsa/main.go`


# into mysql

```
docker compose exec db bash
root@51e66cf1c1d3:/# mysql -uroot -proot
```

## users

```
mysql> select user from user;
+------------------+
| user             |
+------------------+
| docker           |
| root             |
| mysql.infoschema |
| mysql.session    |
| mysql.sys        |
| root             |
+------------------+
6 rows in set (0.02 sec)
```

```
mysql> show grants for docker@'%';
+------------------------------------------------------------+
| Grants for docker@%                                        |
+------------------------------------------------------------+
| GRANT USAGE ON *.* TO `docker`@`%`                         |
| GRANT ALL PRIVILEGES ON `test\_database`.* TO `docker`@`%` |
+------------------------------------------------------------+
2 rows in set (0.01 sec)

```

## protobuf package

build docker image
```
$ pwd
<path>/sandbox/golang/sandbox # current directory

$ docker build --tag protobuf `pwd`/protobuf/
```

generate code
```
$ ./protobuf/make.sh 
```

exec code 
```
go run cmd/grpc/main.go -c or -s
```

## gcp package

### resourcemanager
```
$ gcloud organizations list      
DISPLAY_NAME                 ID  DIRECTORY_CUSTOMER_ID
xxxx                        yyy                    zzz

$ go run cmd/gcp/main.go -org yyy
```