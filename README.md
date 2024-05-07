# go-rest-template
Go (Golang) API REST Template/Boilerplate with Gin Framework

## 1. Run with Docker

1. **Build**

```shell script
go mod download
go get ./...
make build
docker build . -t order-service
```

2. **Run**

```shell script
docker run -p 3000:3000 order-service
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
