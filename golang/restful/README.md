# restful

Clean Archtecture & レイヤーごとの単体テストを勉強するためのディレクトリ

## API

### Create
```bash
curl -X POST -d '{"firstname": "hoge", "lastname":"aaa"}'  http://localhost:8000/create
```

### Read
```bash
curl http://localhost:8000/customer/{id}
```

### Update
```bash
curl -X PATCH -d '{"name": {"first":"foo"}}'  http://localhost:8000/customer/{id}
```

### Delete 
```bash
curl -X DELETE  http://localhost:8000/customer/{id}
```