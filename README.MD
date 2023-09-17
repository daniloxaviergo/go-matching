# go-matching

build
```
docker build -t go-matching .
```

run
```
docker run --rm -p 8080:8080 -it go-matching go run main.go
```

test
```
docker run --rm -p 8080:8080 -it go-matching go test ./...
```

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
