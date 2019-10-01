#!/bin/bash

docker build -t $SERVICE_TAG:$TRAVIS_TAG .

echo "$DOCKER_PASSWORD" | docker login --username $DOCKER_LOGIN --password-stdin;

docker tag $SERVICE_TAG:$TRAVIS_TAG $DOCKER_LOGIN/$SERVICE_TAG:$TRAVIS_TAG;

docker push $DOCKER_LOGIN/$SERVICE_TAG:$TRAVIS_TAG;