BINARY_NAME=goand

build:
	@go mod vendor
	@echo "Building creating"
	@go build -o tmp/${BINARY_NAME}
	@echo "Goand build!"

run: build
	@echo "Starting..."
	@./tmp/${BINARY_NAME} &
	@echo "Goand Started!"

clean:
	@echo "Cleaning.."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

start_compose:
	docker-compose up -d

stop_compose:
	docker-compose down

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