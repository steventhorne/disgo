rwildcard=$(foreach d,$(wildcard $(1:=/*)),$(call rwildcard,$d,$2) $(filter $(subst *,%,$2),$d))
sources := $(call rwildcard,.,*.go)
ifeq ($(OS),Windows_NT)
	bot := bot.exe
	register := register.exe
else
	bot := bot
	register := register
endif

all: $(bot) $(register)

$(bot): $(sources) ./migrations/bindata.go
	go build -o $(bot) ./cmd/bot

$(register): $(sources) ./migrations/bindata.go
	go build -o $(register) ./cmd/register

./migrations/bindata.go: ./migrations/*.sql
	go-bindata -nomemcopy -pkg migrations -ignore bindata -prefix ./migrations/ -o ./migrations/bindata.go ./migrations

.PHONY: clean dep test cover docs

clean:
	go clean

dep:
	go install github.com/kevinburke/go-bindata/v4/...@latest
	go install golang.org/x/tools/cmd/godoc@latest
	go mod download

test:
	go test -v ./...

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

docs:
	godoc -http=:6060 -index -v -play
