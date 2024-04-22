#!/bin/bash

wget https://golang.org/dl/go1.22.1.linux-amd64.tar.gz

tar -C /usr/local -xzf go1.22.1.linux-amd64.tar.gz

export PATH=$PATH:/usr/local/go/bin

go build -o main .

exec "$@"