#!/usr/bin/bash

sudo apt-get install libpam0g-dev libsqlcipher-dev -y
export CGO_ENABLED=1

go build -o mypasswords main.go

sudo cp mypasswords /usr/local/bin/mypasswords