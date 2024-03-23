# Makefile for Equilibria

# Docker image name
IMAGE_NAME=equilibria-api

# Docker container name
CONTAINER_NAME=equilibria-api

AWS_REGION ?= us-west-2
AWS_ACCOUNT_ID ?= 462498369025
IMAGE_NAME := equilibria-api
REPO_NAME := equilibria-api
IMAGE_TAG ?= latest

ECR_URL := $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com


.PHONY: all build format lint lint-readme install-markdownlint docker-build

# Default target
all: build

# Build the project
build:
	go build -o equilibria_api

# Format the Go source files using gofumpt
format:
	go gofumpt -l -w .

# Run golangci-lint
lint:
	golangci-lint run

# Lint the README file using markdownlint
lint-readme: install-markdownlint
	markdownlint README.md

# Install markdownlint
install-markdownlint:
	@echo "Downloading markdownlint..."
	curl -sSLo markdownlint https://github.com/igorshubovych/markdownlint-cli/releases/download/0.27.1/markdownlint
	chmod +x markdownlint

docker-build:
	@echo "Building docker image"
	docker build -t equilibria-api .

docker-run:
	@echo "Running docker container"
    docker run -p 443:443 --env-file .env --name $(CONTAINER_NAME) $(IMAGE_NAME)

docker-up: docker-build docker-run

ecr-auth:
	@echo "Authenticating with AWS ECR..."
	@aws ecr get-login-password --region us-west-2 | docker login --username AWS --password-stdin 462498369025.dkr.ecr.us-west-2.amazonaws.com

deploy: ecr-auth
	# Build and push the Docker image directly to ECR
	@echo "Building and deploying Docker to ECR"
	docker buildx build --platform linux/amd64 -t $(ECR_URL)/$(REPO_NAME):$(IMAGE_TAG) . --push

	@echo "Cycling AWS ECS Cluster"
	aws ecs update-service --cluster eq-api-cluster --service eq-api-service --force-new-deployment

