#!/bin/bash

for i in "$@"
do
case $i in
    -tps=*|--tps=*)
    TPS="${i#*=}"
    shift # past argument=value
    ;;
    -test_size=*|--test_size=*)
    TEST_SIZE="${i#*=}"
    shift # past argument=value
    ;;
    --default)
    DEFAULT=YES
    shift # past argument with no value
    ;;
    *)
          # unknown option
    ;;
esac
done
echo "Building tests"
go build ./go_test_1/transaction.go
go build ./go_test_2/transaction.go
go build ./go_test_3/transaction.go
go build ./go_test_4/transaction.go

echo "Run test size ${TEST_SIZE} with ${TPS} on each node."

nohup ./go_test_2/transaction -tps=${TPS} -test_size=${TEST_SIZE}&
nohup ./go_test_3/transaction -tps=${TPS} -test_size=${TEST_SIZE}&
./go_test_1/transaction -tps=${TPS} -test_size=${TEST_SIZE}
