# Ping-Logger
Simple Golang application that pings a host every X seconds over a set duration. Exits after duration has been reached.

Logs to logfile if ping returns packet loss.

## Config
The config file `config.ini` is assumed to be located in the same folder as the app.
### Config params
| Value | Description | Required | Default |
| ----- | ----------- |:--------:| ------- |
| host | Hostname of server to ping | * | n/a
| interval | Pinging interval in seconds |   | 600 |
| duration | App running duration in seconds |   | 21600 |
| log | Path of logfile |   | ./pinglogger.txt |

## Run
```
go run ping-logger.go
```

## Build
```
go build -ldflags="-X 'main.version=$(git describe --tags --abbrev=0)'"
```
In this case using latest git tag as version number.