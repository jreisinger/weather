.DEFAULT_GOAL := build

test:
	go test

install: test
	go install cmd/weather.go

build: test
	go build cmd/weather.go

run: build
	./weather Bratislava
	./weather Milano
	./weather London