run: build
	@./bin/verx-gm-bot

build:
	@go build -o ./bin/verx-gm-bot ./cmd/verx-gm-bot/main.go

build-cli:
	@go build -o ./bin/verx-tg-cli ./cmd/verx-tg-cli/main.go

build-linux-amd:
	@mkdir -p ./bin/linux/amd64
	@GOOS=linux GOARCH=amd64 go build -o ./bin/linux/amd64/verx-gm-bot ./cmd/verx-gm-bot/main.go