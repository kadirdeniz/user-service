#!/bin/bash

docker run -d --name user-service-redis -p 6379:6379 redis:latest