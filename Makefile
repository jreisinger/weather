test:
	go test

install: test
	go install cmd/weather.go
