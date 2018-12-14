#!/bin/bash

echo "Building tests ..."
cd ./go_test_1/
go build transaction.go
cd ../go_test_2/
go build transaction.go
cd ../go_test_3/
go build transaction.go
cd ../go_test_4/
go build transaction.go
cd ../ask_node/
go build transaction.go
