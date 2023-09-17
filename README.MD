# go-matching

build
```
docker build -t go-matching .
```

run
```
docker run --rm -p 8080:8080 -it go-matching go run main.go
```
![image](https://github.com/daniloxaviergo/go-matching/assets/35167007/b04a2bbe-d649-4069-ab20-76096e80d238)

test
```
docker run --rm -p 8080:8080 -it go-matching go test ./...
```
![image](https://github.com/daniloxaviergo/go-matching/assets/35167007/a7ef2cde-6784-4f04-ae18-f9a0a832f8f3)


request's
```
curl -v 'http://localhost:8080/btcbrl/order-create?id=1&price=12.34&volume=78.85&side=ask'
curl -v 'http://localhost:8080/btcbrl/order-create?id=2&price=12.34&volume=78.85&side=ask'
curl -v 'http://localhost:8080/btcbrl/order-create?id=3&price=21.43&volume=99.05&side=bid'
curl -v 'http://localhost:8080/btcbrl/orderbook'

curl -v 'http://localhost:8080/btcbrl/order-cancel?id=1'
curl -v 'http://localhost:8080/btcbrl/order-cancel?id=2'
curl -v 'http://localhost:8080/btcbrl/orderbook'
```
![image](https://github.com/daniloxaviergo/go-matching/assets/35167007/56989757-4acd-4989-ba9a-570f5fb7b84f)
