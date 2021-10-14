go build main.go
./main serve


docker build .-t pricefeeder:latest


docker run pricefeeder:latest -p  8080:8080
