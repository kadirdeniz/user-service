#!/bin/bash


grep -A3 'mongodb:' ../configs/mongodb.yaml | tail -n1 | awk '{ print $2}'
#mongoimport --drop --host my_mongo --username root --password rootpassword --authenticationDatabase admin --db tutorial_db --collection users --type json --jsonArray --file api/users.json