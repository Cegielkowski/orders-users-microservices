#!/bin/bash

echo "Executing services tests"

cp ./config/tests_config.json.example ./config/tests_config.json

cd pkg/services || exit
go test ./...