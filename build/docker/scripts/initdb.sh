#!/usr/bin/env bash

echo 'Creating application user and db'

mongo ${MONGO_DB} \
        --host localhost \
        --port 27017 \
        -u ${MONGO_ROOT_USER} \
        -p ${MONGO_ROOT_PW} \
        --authenticationDatabase admin \
        --eval "db.createUser({user: '${MONGO_USER}', pwd: '${MONGO_PW}', roles:[{role:'dbOwner', db: '${MONGO_DB}'}]});"
