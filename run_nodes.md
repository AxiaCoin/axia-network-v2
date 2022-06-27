./build/axia \
--genesis=./genesis/localhost_genesis.json \
--http-host=127.0.0.1 \
--http-port=9650 \
--staking-port=9651 \
--db-dir=db/node1 \
--network-id=7890 \
--staking-tls-cert-file=$(pwd)/staking/local/staker1.crt \
--staking-tls-key-file=$(pwd)/staking/local/staker1.key

./build/axia \
--genesis=./genesis/localhost_genesis.json \
--http-host=127.0.0.1 \
--http-port=9652 \
--staking-port=9653 \
--db-dir=db/node2 \
--network-id=7890 \
--bootstrap-ips=127.0.0.1:9651 \
--bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg \
--staking-tls-cert-file=$(pwd)/staking/local/staker2.crt \
--staking-tls-key-file=$(pwd)/staking/local/staker2.key

./build/axia \
--genesis=./genesis/localhost_genesis.json \
--http-host=127.0.0.1 \
--http-port=9654 \
--staking-port=9655 \
--db-dir=db/node3 \
--network-id=7890 \
--bootstrap-ips=127.0.0.1:9651 \
--bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg \
--staking-tls-cert-file=$(pwd)/staking/local/staker3.crt \
--staking-tls-key-file=$(pwd)/staking/local/staker3.key

./build/axia \
--genesis=./genesis/localhost_genesis.json \
--http-host=127.0.0.1 \
--http-port=9656 \
--staking-port=9657 \
--db-dir=db/node4 \
--network-id=7890 \
--bootstrap-ips=127.0.0.1:9651 \
--bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg \
--staking-tls-cert-file=$(pwd)/staking/local/staker4.crt \
--staking-tls-key-file=$(pwd)/staking/local/staker4.key

./axia-network-runner server \
--log-level debug \
--port=":8080" \
--grpc-gateway-port=":8081"


curl --location --request POST '127.0.0.1:9650/ext/bc/Swap' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc":"2.0",
    "id"     :1,
    "method" :"avm.importKey",
    "params" :{
        "username": "sankar",
        "password": "sR-22@met",
        "privateKey":"PrivateKey-ewoqjP7PxY4yr3iLTpLisriqt94hdyDFNgchSxGGztUrTXtNN"
    }
}'

curl --location --request POST '127.0.0.1:9650/ext/bc/Swap' \
--header 'Content-Type: application/json' \
--data-raw '{
  "jsonrpc":"2.0",
  "id"     : 1,
  "method" :"avm.getBalance",
  "params" :{
      "address":"Swap-custom18jma8ppw3nhx5r4ap8clazz0dps7rv5u9xde7p",
      "assetID": "AXC"
  }
}'
