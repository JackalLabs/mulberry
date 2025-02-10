package relay

import (
	"github.com/JackalLabs/mulberry/config"
	"github.com/JackalLabs/mulberry/jackal/uploader"
	"github.com/desmos-labs/cosmos-go-wallet/wallet"
	"github.com/ethereum/go-ethereum/common"
)

type App struct {
	w   *wallet.Wallet
	q   *uploader.Queue
	cfg config.Config
}

var ChainIDS = map[uint64]string{
	1:     "Ethereum",
	8453:  "Base",
	137:   "Polygon",
	10:    "OP",
	42161: "Arbitrum",
}

type PostedFile struct {
	From    common.Address
	Merkle  string
	Size    uint64
	Note    string
	Expires uint64
}

type BoughtStorage struct {
	From         common.Address
	ForAddress   string
	DurationDays uint64
	SizeBytes    uint64
	Referral     string
}

type DeletedFile struct {
	From   common.Address
	Merkle string
	Start  uint64
}
