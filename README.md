
# fib_api

REST API that returns the specified nth Fibonacci number.
If n is a value between 1 and 1000, the nth Fibonacci number is returned.
If n is less than 1 or greater than 1001, the http status 400 Bad request. is returned.

The following is used.
- language: `Go 1.21.0`
- Framework: `Gin`
- hosting: `Heroku`

The entire application is contained within the `main.go` file.
In `fib.go`, the process for Fibonacci numbers is described.
In `fib_test.go`, tests for `fib.go` are written.

In `fib.go`, the Fibonacci numbers are identified using a loop statement.
The process is outlined as adding the first value and the next value together for the number of times num has been counted.
The amount of calculation is O(N).

## Runninge Locally
Make sure you have Go version 1.21 or newer and the Heroku CLI installed.

```shell
$ git clone https://github.com/nanikasi/fib_api.git
$ cd fib_api
$ go build -o bin/fibonacci-api -v .
$ heroku local
```
Your app should now be running on localhost:5000.

## Run the test

```shell
$ go test fib.go fib_test.go
```

## REST API
The REST API to the fibonacci API is described below.
### Request
`GET /fib?n=100`
```shell
curl -i -H 'Accept: application/json' http://localhost:5000/fib?n=100
```
### Response
```shell
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Mon, 16 Oct 2023 00:28:40 GMT
Content-Length: 34

{"result":"354224848179261915075"}
```
