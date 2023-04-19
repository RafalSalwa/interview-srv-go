# Docker Makefile
PROJECT_NAME=interview

.PHONY: build
build:
	docker compose down
	docker compose -p $(PROJECT_NAME) up --build

.PHONY: test
test:
	docker build --target builder -t $(PROJECT_NAME)-ci
	docker run $(PROJECT_NAME)-ci sh -c "npm run test"

.PHONY: clean
clean:
	docker ps -a | awk '/$(PROJECT_NAME)/ { print $$1 }' | xargs docker rm -f
	docker images -a | awk '/$(PROJECT_NAME)/ { print $$3 }' | xargs docker rmi -f