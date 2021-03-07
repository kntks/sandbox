# restful
[![Go](https://github.com/kntks/sandbox/actions/workflows/go.yml/badge.svg)](https://github.com/kntks/sandbox/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/kntks/sandbox/branch/main/graph/badge.svg?token=9ZS9IV1C69)](https://codecov.io/gh/kntks/sandbox)

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