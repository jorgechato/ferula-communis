# Go web service monitored with prometheus
[![Go Report Card](https://goreportcard.com/badge/github.com/jorgechato/ferula-communis)](https://goreportcard.com/report/github.com/jorgechato/ferula-communis)

A demo of a GO web service build with gorilla/mux and monitored with prometheus
## Install & Run
### Web service
```zsh
$ docker build -t ferula-communis .
$ docker run -d -p 8080:8080 --name ferula-communis_web ferula-communis
```

### Set prometheus
Write your own Prometheus configuration to scrape itself. Use the following yaml or crate one yourself:

```yaml
global:
  scrape_interval:     15s
scrape_configs:
  - job_name: 'ferula-communis'
    scrape_interval: 5s
    static_configs:
      - targets: ['ferula-communis_web:8080']
```

Mount that configuration file in a Docker volume and run Prometheus again:

```zsh
$ docker run -p 9090:9090 -v ${PWD}/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus
```

### With docker-compose
```zsh
$ docker-compose up -d
```
