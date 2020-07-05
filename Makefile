APP_NAME=e2e-rest

all: build 

build: clean 
	@echo "----------------------------------------------------------" 
	@echo "Building $(APP_NAME) to bin" 
	@echo "----------------------------------------------------------" 
	@go build -o bin/$(APP_NAME)

format:
	@echo "----------------------------------------------------------" 
	@echo "Formatting code" 
	@echo "----------------------------------------------------------" 
	@go fmt ./...

.PHONY: local-e2e-test
local-e2e-test:
	@echo "----------------------------------------------------------" 
	@echo "Executing tests. Postgre SQL should be started in advance." 
	@echo "----------------------------------------------------------" 
	@go test -v ./...

.PHONY: e2e-test
e2e-test:
	@echo "----------------------------------------------------------" 
	@echo "Executing tests using docker-compose" 
	@echo "----------------------------------------------------------" 
	@echo "----------------------------------------------------------" 
	@echo "Cleaning" 
	@echo "----------------------------------------------------------" 
	@docker-compose -f ./test/docker-compose.yml --env-file ./test/.env down
	@docker-compose -f ./test/docker-compose.yml --env-file ./test/.env rm
	@echo "----------------------------------------------------------" 
	@echo "Building image..." 
	@echo "----------------------------------------------------------" 
	@docker-compose -f ./test/docker-compose.yml --env-file ./test/.env build
	@echo "----------------------------------------------------------" 
	@echo "Composing environment and execute tests..." 
	@echo "----------------------------------------------------------" 
	@docker-compose -f ./test/docker-compose.yml --env-file ./test/.env up --abort-on-container-exit
	@echo "----------------------------------------------------------" 
	@echo "Cleaning" 
	@echo "----------------------------------------------------------" 
	@docker-compose -f ./test/docker-compose.yml --env-file ./test/.env down
	@docker-compose -f ./test/docker-compose.yml --env-file ./test/.env rm
	@echo "----------------------------------------------------------" 
	@echo "Done!" 
	@echo "----------------------------------------------------------" 

clean: 
	@echo "----------------------------------------------------------" 
	@echo "Cleaning" 
	@echo "----------------------------------------------------------" 
	@go clean
	@rm -f results.json
	@rm -f coverage.txt
	@rm -fr bin

docker: 
	@echo "----------------------------------------------------------" 
	@echo "Build docker image" 
	@echo "----------------------------------------------------------" 
	@docker build -t $(APP_NAME) .

docker-nc: 
	@echo "----------------------------------------------------------" 
	@echo "Build docker image without caching" 
	@echo "----------------------------------------------------------" 
	@docker build --no-cache -t $(APP_NAME) .