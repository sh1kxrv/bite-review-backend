#!/bin/bash

go build -ldflags="-s -w" -o ./bin/app .
./bin/app