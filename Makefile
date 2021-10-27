all:
	gofmt -s -w .
	go build -o smeagleasm
	
build:
	go build -o smeagleasm
	
run:
	go run main.go
