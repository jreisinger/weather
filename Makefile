test:
	go test -cover ./...

install: test
	go install ./cmd/weather.go
