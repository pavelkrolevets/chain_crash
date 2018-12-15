#!/bin/bash

source ./congfig.cfg

readonly address_to1=$address_to1
readonly priv_key1=$priv_key1

readonly address_to2=$address_to2
readonly priv_key2=$priv_key2

readonly address_to3=$address_to3
readonly priv_key3=$priv_key3

readonly address_to4=$address_to4
readonly priv_key4=$priv_key4

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
    -rpc_addr=*|--rpc_addr=*)
    RPC_ADDR="${i#*=}"
    shift # past argument=value
    ;;
    -check_tx=*|--check_tx=*)
    CHECK_TX="${i#*=}"
    shift # past argument=value
    ;;
    -threads=*|--threads=*)
    threads="${i#*=}"
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

echo "Run test size ${TEST_SIZE} with ${TPS} on each node. RPC address ${RPC_ADDR}. Checking txs ${CHECK_TX}"

if [ ${threads} == 1 ]
then
    ./go_test_1/transaction -tps=${TPS} -test_size=${TEST_SIZE} -rpc_addr=${RPC_ADDR} -check_txs=${CHECK_TX} -address_to=${address_to1} -priv_key=${priv_key1} > /dev/null 2>&1 &
    ./ask_node/transaction -rpc_addr=${RPC_ADDR}
fi


if [ ${threads} == 2 ]
then
    ./go_test_1/transaction -tps=${TPS} -test_size=${TEST_SIZE} -rpc_addr=${RPC_ADDR} -check_txs=${CHECK_TX} -address_to=${address_to1} -priv_key=${priv_key1} > /dev/null 2>&1 &
    ./go_test_2/transaction -tps=${TPS} -test_size=${TEST_SIZE} -rpc_addr=${RPC_ADDR} -check_txs=${CHECK_TX} -address_to=${address_to2} -priv_key=${priv_key2} > /dev/null 2>&1 &
    ./ask_node/transaction -rpc_addr=${RPC_ADDR}
fi

if [ ${threads} == 3 ]
then
    ./go_test_1/transaction -tps=${TPS} -test_size=${TEST_SIZE} -rpc_addr=${RPC_ADDR} -check_txs=${CHECK_TX} -address_to=${address_to1} -priv_key=${priv_key1} > /dev/null 2>&1 &
    ./go_test_2/transaction -tps=${TPS} -test_size=${TEST_SIZE} -rpc_addr=${RPC_ADDR} -check_txs=${CHECK_TX} -address_to=${address_to2} -priv_key=${priv_key2} > /dev/null 2>&1 &
    ./go_test_3/transaction -tps=${TPS} -test_size=${TEST_SIZE} -rpc_addr=${RPC_ADDR} -check_txs=${CHECK_TX} -address_to=${address_to3} -priv_key=${priv_key3} > /dev/null 2>&1 &
    ./ask_node/transaction -rpc_addr=${RPC_ADDR}
fi

if [ ${threads} == 4 ]
then
    ./go_test_1/transaction -tps=${TPS} -test_size=${TEST_SIZE} -rpc_addr=${RPC_ADDR} -check_txs=${CHECK_TX} -address_to=${address_to1} -priv_key=${priv_key1} > /dev/null 2>&1 &
    ./go_test_2/transaction -tps=${TPS} -test_size=${TEST_SIZE} -rpc_addr=${RPC_ADDR} -check_txs=${CHECK_TX} -address_to=${address_to2} -priv_key=${priv_key2} > /dev/null 2>&1 &
    ./go_test_3/transaction -tps=${TPS} -test_size=${TEST_SIZE} -rpc_addr=${RPC_ADDR} -check_txs=${CHECK_TX} -address_to=${address_to3} -priv_key=${priv_key3} > /dev/null 2>&1 &
    ./go_test_4/transaction -tps=${TPS} -test_size=${TEST_SIZE} -rpc_addr=${RPC_ADDR} -check_txs=${CHECK_TX} -address_to=${address_to4} -priv_key=${priv_key4} > /dev/null 2>&1 &
    ./ask_node/transaction -rpc_addr=${RPC_ADDR}
fi
