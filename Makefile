BINPATH = bin/main


.PHONY: build
build:
		templ generate /internal/view
		go build -o ./bin/main.exe cmd/app/main.go

.PHONY: run
run: build
		./bin/main.exe

.PHONY: watch-app
watch-app:
		go run github.com/air-verse/air@latest \
		--build.bin "$(BINPATH)" \
		--build.include_ext "go" \
		--build.exclude_dir "bin"

.PHONY: watch-templ
watch-templ:
		templ generate --watch --proxy="http://localhost:8080" --open-browser=false

.PHONY: fmt
fmt:
		templ fmt internal/view