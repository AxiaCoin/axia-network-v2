./build/axgo \
--public-ip=127.0.0.1 \
--http-port=9650 \
--staking-port=9651 \
--db-dir=db/node1 \
--network-id=test \
--staking-tls-cert-file=$(pwd)/staking/local/staker1.crt \
--staking-tls-key-file=$(pwd)/staking/local/staker1.key

./build/axgo \
--public-ip=127.0.0.1 \
--http-port=9652 \
--staking-port=9653 \
--db-dir=db/node2 \
--network-id=test \
--bootstrap-ips=127.0.0.1:9651 \
--bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg \
--staking-tls-cert-file=$(pwd)/staking/local/staker2.crt \
--staking-tls-key-file=$(pwd)/staking/local/staker2.key

./build/axgo \
--public-ip=127.0.0.1 \
--http-port=9654 \
--staking-port=9655 \
--db-dir=db/node3 \
--network-id=test \
--bootstrap-ips=127.0.0.1:9651 \
--bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg \
--staking-tls-cert-file=$(pwd)/staking/local/staker3.crt \
--staking-tls-key-file=$(pwd)/staking/local/staker3.key

./build/axgo \
--public-ip=127.0.0.1 \
--http-port=9656 \
--staking-port=9657 \
--db-dir=db/node4 \
--network-id=test \
--bootstrap-ips=127.0.0.1:9651 \
--bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg \
--staking-tls-cert-file=$(pwd)/staking/local/staker4.crt \
--staking-tls-key-file=$(pwd)/staking/local/staker4.key
