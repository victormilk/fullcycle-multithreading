CEP ?= 01153000

run:
	@echo "Running the program..."
	@go run cmd/multithreading/main.go $(CEP)
	@echo "Done!"

PHONY: run
