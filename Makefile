EXEC=api_gateway_loom_chat.bin

build:
	@go build -o $(EXEC) ./cmd

run: build
	@./$(EXEC)

clean:
	@rm ./$(EXEC)

.PHONY:
	run build clean
