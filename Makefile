.DEFAULT_GOAL := build

test:
	go test

install: test
	go install cmd/weather.go
