#!/bin/bash

# Get the short Git SHA
GIT_SHA=$(git rev-parse --short HEAD)

# Docker image name
DOCKER_IMAGE_NAME="gabrie30/hello-world"

# Tag the Docker image
docker tag $DOCKER_IMAGE_NAME $DOCKER_IMAGE_NAME:$GIT_SHA
docker tag $DOCKER_IMAGE_NAME $DOCKER_IMAGE_NAME:latest

# Push the tagged images to Docker Hub
docker push $DOCKER_IMAGE_NAME:$GIT_SHA
docker push $DOCKER_IMAGE_NAME:latest
