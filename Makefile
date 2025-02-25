.PHONY: build
build:
		templ generate /internal/view
		go build -o ./bin/main.exe cmd/app/main.go

.PHONY: run
run: build
		./bin/main.exe