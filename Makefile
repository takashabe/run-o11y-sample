PROJECT_ID   ?=
SERVICE_NAME := run-o11y-sample
REGION       := us-central1

KO_DOCKER_REPO := us-central1-docker.pkg.dev/$(PROJECT_ID)/sandbox
IMAGE_NAME     := run-o11y-sample-9017db5df1611930ee9f148e8c337236

.PHONY: all
all: build deploy

.PHONY: deploy
deploy:
	gcloud run deploy $(SERVICE_NAME) \
		--project $(PROJECT_ID) \
		--image $(KO_DOCKER_REPO)/$(IMAGE_NAME) \
		--platform managed \
		--region $(REGION) \
		--allow-unauthenticated \
		--set-env-vars=PROJECT_ID=$(PROJECT_ID)

.PHONY: build
build:
	KO_DOCKER_REPO=$(KO_DOCKER_REPO) ko publish .
