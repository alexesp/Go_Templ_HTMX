.PHONY: build
build:
		go build -o ./bin/myapp.exe cmd/app/main.go

.PHONY: run
run: build
		./bin/myapp.exe