#!/bin/bash
ip=120.92.123.216
GOOS=linux GOARCH=amd64 go build -o deploy/datalog
scp -r ./deploy/* root@${ip}:/root/golang/