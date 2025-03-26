# base
## process
```bash
export BASE_RPC_URL="https://mainnet.base.org"
export BASE_SEPOLIA_RPC_URL="https://sepolia.base.org"
cast wallet import deployer --interactive # import private key
forge create src/JackalV1.sol:JackalBridge --root forge --rpc-url $BASE_SEPOLIA_RPC_URL --account deployer --constructor-args "[0x190F11f9110aeCDb00a3484E82EB5592B405071D]" 0x4aDC67696bA383F43DD60A9e78F2C97Fbbfc7cb1 # "[relay]" price_feed
forge create src/StorageDrawer.sol:StorageDrawer --root forge --rpc-url $BASE_SEPOLIA_RPC_URL --account deployer --constructor-args 0x6f348699508B317862348f8d6F41795900E8d14A # bridge
# re-run the above two commands with --broadcast
nano ~/.mulberry/config.yaml # add network information
```
## results
```
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0x6f348699508B317862348f8d6F41795900E8d14A
Transaction hash: 0xbda3a5e89cc08d110e6466071d151bfb09ab364aaf84082fb5c70fe5f54c8911
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0x83f69195100eea97BA9Fd0a4e15a1657Efd9D631
Transaction hash: 0xcbdf4ff8bf1b66c3b15ac80c999ddd4a4b48b37d302fbd9ec7ff162c87c8c4eb
```

# optimism
## process
```bash
export OP_RPC_URL="https://mainnet.optimism.io"
export OP_SEPOLIA_RPC_URL="https://sepolia.optimism.io"
cast wallet import deployer --interactive # import private key
forge create src/JackalV1.sol:JackalBridge --root forge --rpc-url $OP_SEPOLIA_RPC_URL --account deployer --constructor-args "[0x190F11f9110aeCDb00a3484E82EB5592B405071D]" 0x61Ec26aA57019C486B10502285c5A3D4A4750AD7 # "[relay]" price_feed
forge create src/StorageDrawer.sol:StorageDrawer --root forge --rpc-url $OP_SEPOLIA_RPC_URL --account deployer --constructor-args 0x82a8d3781241Ab5E5ffF8AB3292765C0f9d0431F # bridge
# re-run the above two commands with --broadcast
nano ~/.mulberry/config.yaml # add network information
```
## results
```
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0x82a8d3781241Ab5E5ffF8AB3292765C0f9d0431F
Transaction hash: 0xd34450a11b1c0c59886039c9c38de5415954f78f8ebc1468b56231f33b28a96b
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0x7dAB0A27c5aB9D1Fb3D2f91E9f0eee9BD051a448
Transaction hash: 0xefe35365c29b4ee42232e49ff29e1bd18e7397f26cbcf02af7c7665d2142a787
```

# ethereum
## process
```bash
export ETH_RPC_URL="https://eth.drpc.org"
export ETH_SEPOLIA_RPC_URL="https://sepolia.drpc.org"
cast wallet import deployer --interactive # import private key
forge create src/JackalV1.sol:JackalBridge --root forge --rpc-url $ETH_SEPOLIA_RPC_URL --account deployer --constructor-args "[0x190F11f9110aeCDb00a3484E82EB5592B405071D]" 0x694AA1769357215DE4FAC081bf1f309aDC325306 # "[relay]" price_feed
forge create src/StorageDrawer.sol:StorageDrawer --root forge --rpc-url $ETH_SEPOLIA_RPC_URL --account deployer --constructor-args 0x1A829964Dd155D89eBA94CfB6CAcbEC496C1df32 # bridge
# re-run the above two commands with --broadcast
nano ~/.mulberry/config.yaml # add network information
```
## results
```
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0x1A829964Dd155D89eBA94CfB6CAcbEC496C1df32
Transaction hash: 0x282266308f62a1b589d1152797be92c0c1bb337bcb095238e3312f45a199baa1
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0xadCAD6Cc46364a6FF0Cb6d5023Af15388C6D17C1
Transaction hash: 0x224f110c2c85c94b3adf45c8f60a862bd34ba18eb7985befa680222525235f37
```

# polygon
## process
```bash
export POLYGON_RPC_URL="https://polygon.drpc.org"
export POLYGON_AMOY_RPC_URL="https://polygon-amoy.drpc.org"
cast wallet import deployer --interactive # import private key
forge create src/JackalV1.sol:JackalBridge --root forge --rpc-url $POLYGON_AMOY_RPC_URL --account deployer --constructor-args "[0x190F11f9110aeCDb00a3484E82EB5592B405071D]" 0xF0d50568e3A7e8259E16663972b11910F89BD8e7 # "[relay]" price_feed
forge create src/StorageDrawer.sol:StorageDrawer --root forge --rpc-url $POLYGON_AMOY_RPC_URL --account deployer --constructor-args 0xc4A028437c4A9e0435771239c31C15fB20eD0274 # bridge
# re-run the above two commands with --broadcast
nano ~/.mulberry/config.yaml # add network information
```
## results
```
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0xc4A028437c4A9e0435771239c31C15fB20eD0274
Transaction hash: 0x3687513902cc57bb7a0b12d94e3e302fd6a53514b65af1a4a18533a2d1a1187c
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0x093BB75ba20F4fe05c31a63ac42B93252C31aE02
Transaction hash: 0x634b949b8b617b29a773c1854cbcded54c40f8a7b20001d729937e7db63c906b
```

# arbitrum
## process
```bash
export ARBITRUM_RPC_URL="https://arb1.arbitrum.io/rpc"
export ARBITRUM_SEPOLIA_RPC_URL="https://sepolia-rollup.arbitrum.io/rpc"
cast wallet import deployer --interactive # import private key
forge create src/JackalV1.sol:JackalBridge --root forge --rpc-url $ARBITRUM_SEPOLIA_RPC_URL --account deployer --constructor-args "[0x190F11f9110aeCDb00a3484E82EB5592B405071D]" 0xd30e2101a97dcbAeBCBC04F14C3f624E67A35165 # "[relay]" price_feed
forge create src/StorageDrawer.sol:StorageDrawer --root forge --rpc-url $ARBITRUM_SEPOLIA_RPC_URL --account deployer --constructor-args 0x82a8d3781241Ab5E5ffF8AB3292765C0f9d0431F # bridge
# re-run the above two commands with --broadcast
nano ~/.mulberry/config.yaml # add network information
```
## results
```
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0x82a8d3781241Ab5E5ffF8AB3292765C0f9d0431F
Transaction hash: 0xb13eb7a9f99dcd1fb11a0e4b1c4f4f2bab26904a73b3feacf12e4b961c5212e0
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0x7dAB0A27c5aB9D1Fb3D2f91E9f0eee9BD051a448
Transaction hash: 0x15523115f9e82bebc1fb642148b5169112c7b45db9baa56b115edc200194f4d7
```

# soneium
## process
```bash
export SONEIUM_RPC_URL="https://rpc.soneium.org"
export SONEIUM_MINATO_RPC_URL="https://rpc.minato.soneium.org"
cast wallet import deployer --interactive # import private key
forge create src/JackalV1.sol:JackalBridge --root forge --rpc-url $SONEIUM_MINATO_RPC_URL --account deployer --constructor-args "[0x190F11f9110aeCDb00a3484E82EB5592B405071D]" 0xCA50964d2Cf6366456a607E5e1DBCE381A8BA807 # "[relay]" price_feed
forge create src/StorageDrawer.sol:StorageDrawer --root forge --rpc-url $SONEIUM_MINATO_RPC_URL --account deployer --constructor-args 0x82a8d3781241Ab5E5ffF8AB3292765C0f9d0431F # bridge
# re-run the above two commands with --broadcast
nano ~/.mulberry/config.yaml # add network information
```
## results
```
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0x82a8d3781241Ab5E5ffF8AB3292765C0f9d0431F
Transaction hash: 0xb6220ebe08b628dd65a5b47bdb71d416b1c367f797cacd65244cc3c6b0a77ba2
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0x7dAB0A27c5aB9D1Fb3D2f91E9f0eee9BD051a448
Transaction hash: 0xe73faa0971d3a1e2d3059eadc3965d1b94585578b990ae652354ff865b589d69
```

# config
The most up-to-date config should be at [config/config.go](config/config.go).
```yaml
jackal_config:
    rpc: https://testnet-rpc.jackalprotocol.com:443
    grpc: jackal-testnet-grpc.polkachu.com:17590
    seed_file: seed.json
    contract: jkl1znt8edwvfpfhhsmx5c406g4y3hk9jh6n6hyca2mhcdaft47jrf0satwq8t
networks_config:
    - name: Sepolia
      rpc: https://ethereum-sepolia-rpc.publicnode.com
      ws: wss://ethereum-sepolia-rpc.publicnode.com
      contract: 0x1A829964Dd155D89eBA94CfB6CAcbEC496C1df32
      chain_id: 11155111
      finality: 2
    - name: Base Sepolia
      rpc: https://base-sepolia-rpc.publicnode.com
      ws: wss://base-sepolia-rpc.publicnode.com
      contract: 0x6f348699508B317862348f8d6F41795900E8d14A
      chain_id: 84532
      finality: 2
    - name: OP Sepolia
      rpc: https://optimism-sepolia-rpc.publicnode.com
      ws: wss://optimism-sepolia-rpc.publicnode.com
      contract: 0x82a8d3781241Ab5E5ffF8AB3292765C0f9d0431F
      chain_id: 11155420
      finality: 2
    - name: Polygon Amoy
      rpc: https://polygon-amoy-bor-rpc.publicnode.com
      ws: wss://polygon-amoy-bor-rpc.publicnode.com
      contract: 0xc4A028437c4A9e0435771239c31C15fB20eD0274
      chain_id: 80002
      finality: 2
    - name: Arbitrum Sepolia
      rpc: https://arbitrum-sepolia-rpc.publicnode.com
      ws: wss://arbitrum-sepolia-rpc.publicnode.com
      contract: 0x82a8d3781241Ab5E5ffF8AB3292765C0f9d0431F
      chain_id: 421614
      finality: 2
    - name: Soneium Minato
      rpc: https://soneium-sepolia-rpc.publicnode.com
      ws: wss://soneium-sepolia-rpc.publicnode.com
      contract: 0x82a8d3781241Ab5E5ffF8AB3292765C0f9d0431F
      chain_id: 1946
      finality: 2
```
