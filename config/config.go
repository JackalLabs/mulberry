package config

import (
	_ "github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

type Config struct {
	JackalConfig   JackalConfig    `yaml:"jackal_config" mapstructure:"jackal_config"`
	NetworksConfig []NetworkConfig `yaml:"networks_config" mapstructure:"networks_config"`
}

type JackalConfig struct {
	RPC      string `yaml:"rpc" mapstructure:"rpc"`
	GRPC     string `yaml:"grpc" mapstructure:"grpc"`
	SeedFile string `yaml:"seed_file" mapstructure:"seed_file"`
	Contract string `yaml:"contract" mapstructure:"contract"`
}

type NetworkConfig struct {
	Name     string `yaml:"name" mapstructure:"name"`
	RPC      string `yaml:"rpc" mapstructure:"rpc"`
	WS       string `yaml:"ws" mapstructure:"ws"`
	Contract string `yaml:"contract" mapstructure:"contract"`
	ChainID  uint64 `yaml:"chain_id" mapstructure:"chain_id"`
	Finality uint64 `yaml:"finality" mapstructure:"finality"`
}

func DefaultConfig() Config {
	return Config{
		JackalConfig: JackalConfig{
			RPC:      "https://testnet-rpc.jackalprotocol.com:443",
			GRPC:     "jackal-testnet-grpc.polkachu.com:17590",
			SeedFile: "seed.json",
			Contract: "jkl163jzm5mmy29ke6ecn4m5evqk9lfhxwne3vzjpe3lpyc33yz5uy0skhhxm8",
		},
		NetworksConfig: []NetworkConfig{
			{
				Name:     "Sepolia",
				RPC:      "https://ethereum-sepolia-rpc.publicnode.com",
				WS:       "wss://ethereum-sepolia-rpc.publicnode.com",
				Contract: "0x093BB75ba20F4fe05c31a63ac42B93252C31aE02",
				ChainID:  11155111,
				Finality: 2,
			},
			{
				Name:     "Base Sepolia",
				RPC:      "https://base-sepolia-rpc.publicnode.com",
				WS:       "wss://base-sepolia-rpc.publicnode.com",
				Contract: "0x5d26f092717A538B446A301C2121D6C68157467C",
				ChainID:  84532,
				Finality: 2,
			},
			{
				Name:     "OP Sepolia",
				RPC:      "https://optimism-sepolia-rpc.publicnode.com",
				WS:       "wss://optimism-sepolia-rpc.publicnode.com",
				Contract: "0xA3FF0a3e8edCd1c1BefBa6e48e847DB9feF82CA2",
				ChainID:  11155420,
				Finality: 2,
			},
			{
				Name:     "Polygon Amoy",
				RPC:      "https://polygon-amoy-bor-rpc.publicnode.com",
				WS:       "wss://polygon-amoy-bor-rpc.publicnode.com",
				Contract: "0x5d26f092717A538B446A301C2121D6C68157467C",
				ChainID:  80002,
				Finality: 2,
			},
			{
				Name:     "Arbitrum Sepolia",
				RPC:      "https://arbitrum-sepolia-rpc.publicnode.com",
				WS:       "wss://arbitrum-sepolia-rpc.publicnode.com",
				Contract: "0xA3FF0a3e8edCd1c1BefBa6e48e847DB9feF82CA2",
				ChainID:  421614,
				Finality: 2,
			},
			{
				Name:     "Soneium Minato",
				RPC:      "https://soneium-sepolia-rpc.publicnode.com",
				WS:       "wss://soneium-sepolia-rpc.publicnode.com",
				Contract: "0xA3FF0a3e8edCd1c1BefBa6e48e847DB9feF82CA2",
				ChainID:  1946,
				Finality: 2,
			},
		},
	}
}

func (c Config) Export() ([]byte, error) {
	return yaml.Marshal(c)
}
