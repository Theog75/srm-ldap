#!/bin/bash

go build .

docker rmi quay.io/ypery/srm-ldap:v$TAG
docker build -t quay.io/ypery/srm-ldap:v$TAG .
docker push quay.io/ypery/srm-ldap:v$TAG


