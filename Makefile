# the name of the application
NAME=dicebot

# GCP Project ID
GCR_TAG=gcr.io/${PROJECT_ID}/$(NAME)

# Container Version
VERSION=v1


## Run the goagen command to generate application files from DSL.
goagen:
	goagen bootstrap -d $(NAME)/design

## Run the docker build command to create the container and compile the
## application inside the container.
build:
	docker build -t $(GCR_TAG):$(VERSION) .

## Run the test suite inside the docker container
test: build
	docker run --name $(NAME)-test --rm $(NAME) go test -v ./...

## Build and run the application using ENV_FILE for configuration.
run: cleanup build
	docker run -it -p 8080:8080 --name $(NAME) $(GCR_TAG):$(VERSION)

## Clean up the existing container reference if it exists
cleanup:
	docker rm $(NAME)

## Publish container to GCP Repository
push:
	gcloud docker -- push $(GCR_TAG):$(VERSION)

## Create a Kubernetes pod for deploying the container to
createpod:
	kubectl run $(NAME) --image=$(GCR_TAG):$(VERSION) --port=8080

