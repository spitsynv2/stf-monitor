#!/bin/sh
VERSION="1.0"

docker build --platform=linux/amd64 --no-cache -f deploy/Dockerfile -t stf-monitor:$VERSION .
