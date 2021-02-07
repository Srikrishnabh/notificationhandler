build:
	go mod tidy -v
	go mod vendor
	go build -o producer cmd/notificationproducer/main.go
	go build -o consumer cmd/notificationhandler/main.go
	docker build . --network=host -f Dockerfile.producer -t producer:1.0
	docker build . --network=host -f Dockerfile.consumer -t consumer:1.0
