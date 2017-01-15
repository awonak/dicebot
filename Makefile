# the name of the application
NAME=dicebot

## Run the docker build command to create the container and compile the
## application inside the container.
build:
	docker build -t $(NAME) --rm .

## Build and run the application using ENV_FILE for configuration.
run: build
	docker run -it --publish=8080:8080 --name $(NAME)-dev --rm $(NAME)

## Run the test suite inside the docker container
test: build
	docker run -e "GIN_MODE=test" --name $(NAME)-test --rm $(NAME) go test -v -race ./...

## Run the goagen command to generate application files from DSL.
goagen:
	goagen bootstrap -d $(NAME)/design

## Generate files from DSL and build the container.
buildgen: goagen build

## Clean up the existing container reference if it exists
cleanup:
	docker rm $(NAME)-dev
