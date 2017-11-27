# Go web service monitored with prometheus
A demo of a GO web service build with gorilla/mux and monitored with prometheus
## Install & Run
```zsh
$ docker build -t ferula-communis .
$ docker run -p 8080:8080 --name ferula-communis_web ferula-communis
```
