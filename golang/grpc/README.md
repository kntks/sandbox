# start server
```
$ docker compose up
```

```
$  curl -X POST -k http://localhost:7090/v1/example/echo -d '{"name": " hello"}'
```

# buf
```
$ cd protobuf
$ buf generate 
```