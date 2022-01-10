# grafana

# Usage
```
$ cp env-template .env
$ docker-compose up
```

access: http://localhost:3000

## into container
```
$ docker-compose exec grafana /bin/bash
bash-5.1$ 
```

## reload config
```
$ docker restart grafana
```
https://grafana.com/docs/grafana/latest/installation/restart-grafana/”