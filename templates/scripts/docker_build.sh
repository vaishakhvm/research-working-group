#! /bin/sh
echo Build Docker Image && docker build -t go/hello . && echo Docker Image Created && ls | grep docker images;
