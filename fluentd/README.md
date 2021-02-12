# fluentd-sandbox
dockerを利用したfluentd勉強用のサンドボックス

https://github.com/fluent/fluentd-docker-image#providing-your-own-configuration-file-and-additional-options

## Usage
### docker compose
 
```bash
cp env-template .env
```

```bash
docker-compose up
```

how to switch config file?

```bash
# .env
# you can choose file from fluentd/fluentd/etc
FLUENTD_CONF=fluent.conf
```

### docker cli
default: fluentd/etc/fluent.conf

```bash
docker run -ti --rm -p 24224:24224 --name fluentd-sandbox --mount type=bind,source=`pwd`/fluentd,target=/fluentd fluentd  -c /fluentd/etc/fluent.conf
```
https://docs.fluentd.org/configuration/config-file#docker

### Into a running container
```bash
# docker-compose
docker-compose exec fluentd  /bin/sh 

# docker cli
docker exec -it fluentd-sandbox /bin/sh
```

## /bin/entrypoint.sh failed: Permission denied

```
chmod +x entrypoint.sh
docker-compose build --no-cache
docker-compose up
```

## Example

### fluentd/etc/intput-http.conf
input
```bash
curl -X POST -d @data/input1.log  http://localhost:24224

# or 

curl -X POST -d 'json={"test":"data"}' http://localhost:24224
```

output
```bash
fluentd_1  | 2021-02-12 02:43:01.162710800 +0000 : {"test":"data"}
```

### fluentd/etc/output-file.conf

input
```bash
curl -X POST -d  @data/input1.log  http://localhost:24224
```

output  

fluentd/log/buffer.xxxx.log
```log
2021-01-16T04:56:31+00:00		{"test":"data"}
```

### fluentd/etc/input-tail.conf

data/input2.log
```
{"test": "hoge", "name": "xxxxx", "age": 20}
```

input
```bash
cat data/input2.log >> fluentd/log/tail.log
```

output
```bash
2021-01-16T05:21:08+00:00	tail.log	{"test":"hoge","name":"xxxxx","age":20}
```

### fluentd/etc/filter-record_transformer.conf

data/input3.log
```
json={"test": "hoge", "name": "xxxxx", "age": 20}
```

input
```bash
curl -X POST -d  @data/input3.log  http://localhost:24224
```

output
```bash
2021-01-16 06:03:40.621452069 +0000 : {"test":"hoge"}
```

### fluentd/etc/out-cloudwatch-logs.conf

set `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`  

input
```bash
cat data/input2.log >> fluentd/log/tail.log
```

output
```bash
2021-01-16 15:45:12.897206945 +0000 tail.log: {"test":"hoge","name":"xxxxx","age":20}
```

AWS console  
CloudWatch > CloudWatch Logs > Log groups > test-log-group-name

## Plugin

### list plugin
```bash
fluent-gem list
```

[fluentd plugin list](https://www.fluentd.org/plugins/all)