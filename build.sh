#!/bin/bash

go build .

docker rmi srm-ldap
docker build -t quay.io/ypery/srm-ldap:v$TAG .
docker push quay.io/ypery/srm-ldap:v$TAG


