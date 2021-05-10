#!/bin/bash

go build .

docker rmi srm-dhcp
docker build -t quay.io/ypery/srm-dhcp:v$TAG .
docker push quay.io/ypery/srm-dhcp:v$TAG


