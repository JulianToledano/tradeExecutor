PROJECT=tradeExecutor


test:
	go test -race -cover $(PROJECT)/...

mod:
	@cd $(PROJECT);	pwd; go mod tidy -v; go mod download;

build:
	@cd $(PROJECT);	pwd; go mod tidy -v; go mod download; go build -o ../bin/tradeExecutor main.go

run:
	go run $(PROJECT)/main.go

clean:
	rm bin/*; rmdir bin