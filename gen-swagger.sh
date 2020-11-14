#!/bin/bash


if [[ $(sudo docker images -q quay.io/goswagger/swagger) == "" ]]; then
    sudo docker pull quay.io/goswagger/swagger
fi

sudo docker run --rm -it -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger generate spec -o $(pwd)/api/swaggerui/swagger.json --scan-models