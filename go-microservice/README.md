# Go Microservice Template

This is a template for creating a microservice in go.

It illustrates a number of features relevant to creating a production-ready microservices in go.

## Features

1. How to create a small modular mircroservice.
2. How to get configurations from CLI arguments.
3. How to use a custom logger.
4. How to do dependency injection.
5. How to create a secure web server using https.
6. How to follow best practices for the TLS configuraitons.
7. How to set content type to speed up the requests.
8. How to create own server to have better control on timeouts.
9. How to add middlewares.
10. How to use docker multi-stage builds to build and ship the service.

## Testing & Running

### 1. Clone & Run

~~~~
- git clone https://github.com/wkabbani/go-samples
- cd go-samples/go-microservice
- go run .
~~~~

### 2. Test

~~~~
// ignore ssl
curl https://localhost:8080/ -k
// or point to the cert
curl https://localhost:8080/ --cacert client-certs/ca.pem
~~~~

### 3. Docker

~~~~
// build
docker build -t [sometag] .
// run
docker run -p 8080:8080 [sometag]
~~~~


## References
- [GopherCon UK 2018: Florin Patan - Production Ready Go Service in 30 Minutes](https://www.youtube.com/watch?v=wxkEQxvxs3w)
- [Go best practices, six years in](https://peter.bourgon.org/go-best-practices-2016/)
- [So you want to expose Go on the Internet](https://blog.cloudflare.com/exposing-go-on-the-internet/)


