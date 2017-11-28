# Go web service monitored with prometheus
[![Go Report Card](https://goreportcard.com/badge/github.com/jorgechato/ferula-communis)](https://goreportcard.com/report/github.com/jorgechato/ferula-communis)

A demo of a GO web service build with gorilla/mux and monitored with prometheus
## Install & Run
```zsh
$ docker build -t ferula-communis .
$ docker run -d -p 8080:8080 --name ferula-communis_web ferula-communis
```
