# Flaky Server

HTTP server that will inconsistently server 200 and 500 responses.

## Install
```
go get github.com/bengadbois/flakyserver
```

```
Usage of flakyserver:
  -errorPercent float
        Percentage of the requests that should return a server error (default 50)
  -port int
        Which port to listen on (default 8899)
````
