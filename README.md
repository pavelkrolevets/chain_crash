# ETH private chain-crash test

This is test to check maximum speed of a private ehthereum blockchain, built on a geth client.
It uses 4 clients which send specified amount of balance transactions each second between prefunded accounts.

##### 1. Confugure prefunded client accounts in `config.cfg`. Note - it should be prefunded otherwise there will be an error.
 
##### 2. Build go instances `./build.sh` . 
Note: everything should be inside `$GOPATH` to build it successfully.

##### 3. Run test `./run_test.sh` with parameters:

- `-threads=2` how many clients will be run in total. Min 1, Max 4. 
- `-tps=50` specifies the speed tx/sec transactions will be generated by each client.
- `-test_size=3000` how many transactions will be produced in total by each client.
- `-rpc_addr=http://127.0.0.1:8545` address and an open port of a mining node.
- `-check_tx=0` after all transactions are sent by each client check whether they are in the chain by getting all of the receipts. Saves receipts to a txt fine in the root folder. 

Example <br>
 `./run_test.sh -tps=50 -test_size=3000 -rpc_addr=http://127.0.0.1:8545 -check_tx=0 -threads=2`

