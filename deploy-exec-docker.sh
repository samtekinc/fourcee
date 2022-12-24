#!/bin/bash

docker build -t tfom-exec-service-executor -f docker/executor.Dockerfile --platform linux/amd64 .

REGION="us-east-1"
REPO_URL=$(aws sts get-caller-identity --output text --query Account).dkr.ecr.$REGION.amazonaws.com/tfom-exec-service-executor
aws ecr get-login-password --region=$REGION |  docker login --username AWS --password-stdin $(aws sts get-caller-identity --output text --query Account).dkr.ecr.$REGION.amazonaws.com
docker tag tfom-exec-service-executor:latest $REPO_URL
docker push $REPO_URL