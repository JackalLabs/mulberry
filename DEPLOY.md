# base
## process
```bash
export BASE_RPC_URL="https://mainnet.base.org"
export BASE_SEPOLIA_RPC_URL="https://sepolia.base.org"
cast wallet import deployer --interactive # import private key
forge create src/JackalV1.sol:JackalBridge --root forge --rpc-url $BASE_SEPOLIA_RPC_URL --account deployer --constructor-args "[0x8792729C879B8B6436e3Dcae8780955ed92F5Af1]" 0x4aDC67696bA383F43DD60A9e78F2C97Fbbfc7cb1 # "[relay]" price_feed
forge create src/StorageDrawer.sol:StorageDrawer --root forge --rpc-url $BASE_SEPOLIA_RPC_URL --account deployer --constructor-args 0x5d26f092717A538B446A301C2121D6C68157467C # bridge
# re-run the above two commands with --broadcast
nano ~/.mulberry/config.yaml # add network information
```
## results
```
[⠊] Compiling...
[⠔] Compiling 8 files with Solc 0.8.28
[⠒] Solc 0.8.28 finished in 395.96ms
Compiler run successful!
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0x5d26f092717A538B446A301C2121D6C68157467C
Transaction hash: 0xdf381d1979a10bf0b0c2629e1e1afef3bb8a4f1efd7f0d06a15588e50a6813d9
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0x82a8d3781241Ab5E5ffF8AB3292765C0f9d0431F
Transaction hash: 0xec9312c8ffc2506f8be85d0c0826338f41df0bafcefa012e151344a20d7fab44
```

# optimism
## process
```bash
export OP_RPC_URL="https://mainnet.optimism.io"
export OP_SEPOLIA_RPC_URL="https://sepolia.optimism.io"
cast wallet import deployer --interactive # import private key
forge create src/JackalV1.sol:JackalBridge --root forge --rpc-url $OP_SEPOLIA_RPC_URL --account deployer --constructor-args "[0x8792729C879B8B6436e3Dcae8780955ed92F5Af1]" 0x61Ec26aA57019C486B10502285c5A3D4A4750AD7 # "[relay]" price_feed
forge create src/StorageDrawer.sol:StorageDrawer --root forge --rpc-url $OP_SEPOLIA_RPC_URL --account deployer --constructor-args 0xA3FF0a3e8edCd1c1BefBa6e48e847DB9feF82CA2 # bridge
# re-run the above two commands with --broadcast
nano ~/.mulberry/config.yaml # add network information
```
## results
```
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0xA3FF0a3e8edCd1c1BefBa6e48e847DB9feF82CA2
Transaction hash: 0x70e220f428428656778ac4614cd379b0a149857c8504d6c13436f6f0a9e50e22
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0x5d26f092717A538B446A301C2121D6C68157467C
Transaction hash: 0x735450af19739ec0fefdd0b419e68691238cbdd999e8b17c05ad85e32aa57e10
```

# ethereum
## process
```bash
export ETH_RPC_URL="https://eth.drpc.org"
export ETH_SEPOLIA_RPC_URL="https://sepolia.drpc.org"
cast wallet import deployer --interactive # import private key
forge create src/JackalV1.sol:JackalBridge --root forge --rpc-url $ETH_SEPOLIA_RPC_URL --account deployer --constructor-args "[0x8792729C879B8B6436e3Dcae8780955ed92F5Af1]" 0x694AA1769357215DE4FAC081bf1f309aDC325306 # "[relay]" price_feed
forge create src/StorageDrawer.sol:StorageDrawer --root forge --rpc-url $ETH_SEPOLIA_RPC_URL --account deployer --constructor-args 0x093BB75ba20F4fe05c31a63ac42B93252C31aE02 # bridge
# re-run the above two commands with --broadcast
nano ~/.mulberry/config.yaml # add network information
```
## results
```
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0x093BB75ba20F4fe05c31a63ac42B93252C31aE02
Transaction hash: 0xf377695c22fb2adeebd11c8d4d4775ab3657bd5e7e2f8ad7dea3d1f6352c932c
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0xB51DC38E8d6EDB08D20B145631F94Ac9d91455d0
Transaction hash: 0xd8c1f5844a4e9bfaac4be2c9de33d73ce058760f3497e50ef0d9d778b7b36132
```

# polygon
## process
```bash
export POLYGON_RPC_URL="https://polygon.drpc.org"
export POLYGON_AMOY_RPC_URL="https://polygon-amoy.drpc.org"
cast wallet import deployer --interactive # import private key
forge create src/JackalV1.sol:JackalBridge --root forge --rpc-url $POLYGON_AMOY_RPC_URL --account deployer --constructor-args "[0x8792729C879B8B6436e3Dcae8780955ed92F5Af1]" 0xF0d50568e3A7e8259E16663972b11910F89BD8e7 # "[relay]" price_feed
forge create src/StorageDrawer.sol:StorageDrawer --root forge --rpc-url $POLYGON_AMOY_RPC_URL --account deployer --constructor-args 0x5d26f092717A538B446A301C2121D6C68157467C # bridge
# re-run the above two commands with --broadcast
nano ~/.mulberry/config.yaml # add network information
```
## results
```
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0x5d26f092717A538B446A301C2121D6C68157467C
Transaction hash: 0x119f426f6c95df2ba7dc8d61189556c82e6441082508f0309723ee8befa378a6
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0x7dAB0A27c5aB9D1Fb3D2f91E9f0eee9BD051a448
Transaction hash: 0x2437fa2e7af22c308e09ff3e2d6d06f403349675c5ecf52b6eca7074097a6785
```

# arbitrum
## process
```bash
export ARBITRUM_RPC_URL="https://arb1.arbitrum.io/rpc"
export ARBITRUM_SEPOLIA_RPC_URL="https://sepolia-rollup.arbitrum.io/rpc"
cast wallet import deployer --interactive # import private key
forge create src/JackalV1.sol:JackalBridge --root forge --rpc-url $ARBITRUM_SEPOLIA_RPC_URL --account deployer --constructor-args "[0x8792729C879B8B6436e3Dcae8780955ed92F5Af1]" 0xd30e2101a97dcbAeBCBC04F14C3f624E67A35165 # "[relay]" price_feed
forge create src/StorageDrawer.sol:StorageDrawer --root forge --rpc-url $ARBITRUM_SEPOLIA_RPC_URL --account deployer --constructor-args 0xA3FF0a3e8edCd1c1BefBa6e48e847DB9feF82CA2 # bridge
# re-run the above two commands with --broadcast
nano ~/.mulberry/config.yaml # add network information
```
## results
```
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0xA3FF0a3e8edCd1c1BefBa6e48e847DB9feF82CA2
Transaction hash: 0x52930c12b158e02a5a1bd251d0c1cc951a85fdc68f592348d611224cf7937046
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0x5d26f092717A538B446A301C2121D6C68157467C
Transaction hash: 0x42eed98b7267085a8764b10dc6804af17dfd59682d96b1a770b43713f2758cd5
```

# soneium
## process
```bash
export SONEIUM_RPC_URL="https://rpc.soneium.org"
export SONEIUM_MINATO_RPC_URL="https://rpc.minato.soneium.org"
cast wallet import deployer --interactive # import private key
forge create src/JackalV1.sol:JackalBridge --root forge --rpc-url $SONEIUM_MINATO_RPC_URL --account deployer --constructor-args "[0x8792729C879B8B6436e3Dcae8780955ed92F5Af1]" 0xCA50964d2Cf6366456a607E5e1DBCE381A8BA807 # "[relay]" price_feed
forge create src/StorageDrawer.sol:StorageDrawer --root forge --rpc-url $SONEIUM_MINATO_RPC_URL --account deployer --constructor-args 0xA3FF0a3e8edCd1c1BefBa6e48e847DB9feF82CA2 # bridge
# re-run the above two commands with --broadcast
nano ~/.mulberry/config.yaml # add network information
```
## results
```
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0xA3FF0a3e8edCd1c1BefBa6e48e847DB9feF82CA2
Transaction hash: 0x65691983fb3c1f0c778fe505d8bab30209d8847d33e3d554e6d7e92e6df54561
[⠊] Compiling...
No files changed, compilation skipped
Enter keystore password:
Deployer: 0x8792729C879B8B6436e3Dcae8780955ed92F5Af1
Deployed to: 0x5d26f092717A538B446A301C2121D6C68157467C
Transaction hash: 0x8f5c1256d77e498dd3e100f723eb56c0f86c1019b0aa72241b15393c81f67081
```

# config
The most up-to-date config should be at [config/config.go](config/config.go).
```yaml
jackal_config:
    rpc: https://testnet-rpc.jackalprotocol.com:443
    grpc: jackal-testnet-grpc.polkachu.com:17590
    seed_file: seed.json
    contract: jkl163jzm5mmy29ke6ecn4m5evqk9lfhxwne3vzjpe3lpyc33yz5uy0skhhxm8
networks_config:
    - name: Sepolia
      rpc: https://ethereum-sepolia-rpc.publicnode.com
      ws: wss://ethereum-sepolia-rpc.publicnode.com
      contract: 0x093BB75ba20F4fe05c31a63ac42B93252C31aE02
      chain_id: 11155111
      finality: 2
    - name: Base Sepolia
      rpc: https://base-sepolia-rpc.publicnode.com
      ws: wss://base-sepolia-rpc.publicnode.com
      contract: 0x5d26f092717A538B446A301C2121D6C68157467C
      chain_id: 84532
      finality: 2
    - name: OP Sepolia
      rpc: https://optimism-sepolia-rpc.publicnode.com
      ws: wss://optimism-sepolia-rpc.publicnode.com
      contract: 0xA3FF0a3e8edCd1c1BefBa6e48e847DB9feF82CA2
      chain_id: 11155420
      finality: 2
    - name: Polygon Amoy
      rpc: https://polygon-amoy-bor-rpc.publicnode.com
      ws: wss://polygon-amoy-bor-rpc.publicnode.com
      contract: 0x5d26f092717A538B446A301C2121D6C68157467C
      chain_id: 80002
      finality: 2
    - name: Arbitrum Sepolia
      rpc: https://arbitrum-sepolia-rpc.publicnode.com
      ws: wss://arbitrum-sepolia-rpc.publicnode.com
      contract: 0xA3FF0a3e8edCd1c1BefBa6e48e847DB9feF82CA2
      chain_id: 421614
      finality: 2
    - name: Soneium Minato
      rpc: https://soneium-sepolia-rpc.publicnode.com
      ws: wss://soneium-sepolia-rpc.publicnode.com
      contract: 0xA3FF0a3e8edCd1c1BefBa6e48e847DB9feF82CA2
      chain_id: 1946
      finality: 2
```
