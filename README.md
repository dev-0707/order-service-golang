# go-rest-template
Go (Golang) API REST Template/Boilerplate with Gin Framework

[![Go Report Card](https://goreportcard.com/badge/order-service)](https://goreportcard.com/report/order-service)
[![Open Source Love](https://badges.frapsoft.com/os/mit/mit.svg?v=102)](https://github.com/ellerbrock/open-source-badge/)
[![Build Status](https://travis-ci.com/antonioalfa22/go-rest-template.svg?branch=master)](https://travis-ci.com/antonioalfa22/go-rest-template)


## 1. Run with Docker

1. **Build**

```shell script
go mod download
go get ./...
make build
docker build . -t api-rest
```

2. **Run**

```shell script
docker run -p 3000:3000 api-rest
```

3. **Test**

```shell script
go test -v ./test/...
```

_______

## 2. Generate Docs

```shell script
# Get swag
go get -u github.com/swaggo/swag/cmd/swag


```

# Generate docs
```
swag init --dir cmd/api --parseDependency --output docs
```


# Test REST API
```
curl "http://localhost:3000/order-service/api/v1/orders/1
```


```
Run and go to **http://localhost:3000/docs/index.html**
```
