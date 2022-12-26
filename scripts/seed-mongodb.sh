#!/bin/bash

host=`grep 'host:' ../configs/mongodb.yaml | tail -n1 | awk '{ print $2}'`
password=`grep 'password:' ../configs/mongodb.yaml | tail -n1 | awk '{ print $2}'`
username=`grep 'username:' ../configs/mongodb.yaml | tail -n1 | awk '{ print $2}'`
port=`grep 'port:' ../configs/mongodb.yaml | tail -n1 | awk '{ print $2}'`
database=`grep 'database:' ../configs/mongodb.yaml | tail -n1 | awk '{ print $2}'`
collection=`grep 'collection:' ../configs/mongodb.yaml | tail -n1 | awk '{ print $2}'`

mongoimport  --drop --host $host --username $username --password $password --authenticationDatabase admin --db $database --collection $collection  --type json  --jsonArray  --file ../api/users.json