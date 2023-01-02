#!/bin/bash

cd configs

host=`grep 'host:' ../mongodb.yaml | tail -n1 | awk '{ print $2}'`
password=`grep 'password:' ../mongodb.yaml | tail -n1 | awk '{ print $2}'`
username=`grep 'username:' ../mongodb.yaml | tail -n1 | awk '{ print $2}'`
port=`grep 'port:' ../mongodb.yaml | tail -n1 | awk '{ print $2}'`
database=`grep 'database:' ../mongodb.yaml | tail -n1 | awk '{ print $2}'`
collection=`grep 'collection:' ../mongodb.yaml | tail -n1 | awk '{ print $2}'`

mongoimport  --drop --host $host --username $username --password $password --authenticationDatabase admin --db $database --collection $collection  --type json  --jsonArray  --file ../api/users.json