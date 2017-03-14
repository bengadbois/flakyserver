# Flaky Server [![Travis](https://img.shields.io/travis/bengadbois/flakyserver.svg?branch=master&style=flat-square)](https://travis-ci.org/bengadbois/flakyserver)

HTTP server that will inconsistently respond with 200 and 500 responses.

## Install
```
go get github.com/bengadbois/flakyserver
```

```
Usage of flakyserver:
  -e float
        Percentage of the requests that should return a server error (default 50)
  -p int
        Which port to listen on (default 8899)
  -t int
        Maximum amount of time to randomly stall (milliseconds)
````
