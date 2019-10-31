ADDR := 127.0.0.1
PORT := 8080

build:
	go build -o rv-host-agent main.go

run: build
	./rv-host-agent server --host $(ADDR) --port $(PORT)
