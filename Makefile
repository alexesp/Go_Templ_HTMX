.PHONY: build
build:
		go build -o ./bin/main.exe cmd/app/main.go

.PHONY: run
run: build
		./bin/main.exe