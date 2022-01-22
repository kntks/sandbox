# grafana

# Usage
```
$ cp env-template .env
$ docker-compose up
```

access: http://localhost:3000

## into container
grafana
```
$ docker compose exec grafana /bin/bash
bash-5.1$ 
```

loki
```
$ docker compose exec loki /bin/ash 
```

promtail
```
$ docker compose exec promtail /bin/bash
```

## reload config
```
$ docker restart grafana
```
https://grafana.com/docs/grafana/latest/installation/restart-grafana/”