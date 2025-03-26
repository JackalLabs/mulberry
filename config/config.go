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
			Contract: "jkl1znt8edwvfpfhhsmx5c406g4y3hk9jh6n6hyca2mhcdaft47jrf0satwq8t",
		},
		NetworksConfig: []NetworkConfig{
			{
				Name:     "Sepolia",
				RPC:      "https://ethereum-sepolia-rpc.publicnode.com",
				WS:       "wss://ethereum-sepolia-rpc.publicnode.com",
				Contract: "0x1A829964Dd155D89eBA94CfB6CAcbEC496C1df32",
				ChainID:  11155111,
				Finality: 2,
			},
			{
				Name:     "Base Sepolia",
				RPC:      "https://base-sepolia-rpc.publicnode.com",
				WS:       "wss://base-sepolia-rpc.publicnode.com",
				Contract: "0x6f348699508B317862348f8d6F41795900E8d14A",
				ChainID:  84532,
				Finality: 2,
			},
			{
				Name:     "OP Sepolia",
				RPC:      "https://optimism-sepolia-rpc.publicnode.com",
				WS:       "wss://optimism-sepolia-rpc.publicnode.com",
				Contract: "0x82a8d3781241Ab5E5ffF8AB3292765C0f9d0431F",
				ChainID:  11155420,
				Finality: 2,
			},
			{
				Name:     "Polygon Amoy",
				RPC:      "https://polygon-amoy-bor-rpc.publicnode.com",
				WS:       "wss://polygon-amoy-bor-rpc.publicnode.com",
				Contract: "0xc4A028437c4A9e0435771239c31C15fB20eD0274",
				ChainID:  80002,
				Finality: 2,
			},
			{
				Name:     "Arbitrum Sepolia",
				RPC:      "https://arbitrum-sepolia-rpc.publicnode.com",
				WS:       "wss://arbitrum-sepolia-rpc.publicnode.com",
				Contract: "0x82a8d3781241Ab5E5ffF8AB3292765C0f9d0431F",
				ChainID:  421614,
				Finality: 2,
			},
			{
				Name:     "Soneium Minato",
				RPC:      "https://soneium-sepolia-rpc.publicnode.com",
				WS:       "wss://soneium-sepolia-rpc.publicnode.com",
				Contract: "0x82a8d3781241Ab5E5ffF8AB3292765C0f9d0431F",
				ChainID:  1946,
				Finality: 2,
			},
		},
	}
}

func (c Config) Export() ([]byte, error) {
	return yaml.Marshal(c)
}
