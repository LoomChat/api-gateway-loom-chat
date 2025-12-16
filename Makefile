EXEC=api_gateway_loom_chat.bin

build:
	@go build -o $(EXEC) ./cmd

run: build
	@./$(EXEC)

format:
	@go fmt ./...

clean:
	@rm ./$(EXEC)

.PHONY:
	run build clean
