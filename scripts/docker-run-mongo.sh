#!/bin/bash

cd configs

host=`grep 'host:' mongodb.yaml | tail -n1 | awk '{ print $2}'`
password=`grep 'password:' mongodb.yaml | tail -n1 | awk '{ print $2}'`
username=`grep 'username:' mongodb.yaml | tail -n1 | awk '{ print $2}'`
port=`grep 'port:' mongodb.yaml | tail -n1 | awk '{ print $2}'`
database=`grep 'database:' mongodb.yaml | tail -n1 | awk '{ print $2}'`
collection=`grep 'collection:' mongodb.yaml | tail -n1 | awk '{ print $2}'`

docker run --name user-service-mongo -p 27017:27017 -e MONGO_INITDB_DATABASE=$database -e MONGO_INITDB_ROOT_USERNAME=$username -e MONGO_INITDB_ROOT_PASSWORD=$password -d mongo