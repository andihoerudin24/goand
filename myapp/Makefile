BINARY_NAME=goand

build:
	@go mod vendor
	@echo "Building creating"
	@go build -o tmp/${BINARY_NAME}
	@echo "Celeritas built!"

run: build
	@echo "Starting..."
	@./tmp/${BINARY_NAME} &
	@echo "Goand Started!"

clean:
	@echo "Cleaning.."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "stopping goand"
	@-pkill SIGTERM -f "./tmp/tmp/${BINARY_NAME}"
	@echo  "Stopped goand"

restart: stop start