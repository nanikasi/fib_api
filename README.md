
# fibonacci-api

A Fibonacci API, which can return the first to 1000th terms of the Fibonacci sequence.

The entire application is contained within the main.go file.

## Runninge Locally
Make sure you have Go version 1.21 or newer and the Heroku CLI installed.
```
$ git clone https://github.com/nanikasi/fibonacci-api.git
$ cd fibonacci-api
$ go build -o bin/fibonacci-api -v . # or `go build -o bin/fibonacci-api.exe -v .` in git bash
$ go start bin/fibonacci-api
```
Your app should now be running on localhost:5000.

## Run the test
```
$ go test fib.go fib_test.go
```

## REST API
The REST API to the fibonacci API is described below.
### Request
`GET /fib?n=100`
```
curl -i -H 'Accept: application/json' http://localhost:5000/fib?n=100
```
### Response
```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Mon, 16 Oct 2023 00:28:40 GMT
Content-Length: 34

{"result":"354224848179261915075"}
```
