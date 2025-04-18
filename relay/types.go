package relay

import (
	"github.com/JackalLabs/mulberry/config"
	"github.com/JackalLabs/mulberry/jackal/uploader"
	"github.com/desmos-labs/cosmos-go-wallet/wallet"
	"github.com/ethereum/go-ethereum/common"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type App struct {
	w    *wallet.Wallet
	wEth *hdwallet.Wallet
	q    *uploader.Queue
	cfg  config.Config
}

var ChainIDS = map[uint64]string{
	1:     "Ethereum",
	8453:  "Base",
	137:   "Polygon",
	10:    "OP",
	42161: "Arbitrum",
	1868:  "Soneium",
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

type RequestedReportForm struct {
	From   common.Address
	Prover string
	Merkle string
	Owner  string
	Start  uint64
}

type PostedKey struct {
	From common.Address
	Key  string
}

type DeletedFileTree struct {
	From     common.Address
	HashPath string
	Account  string
}

type ProvisionedFileTree struct {
	From           common.Address
	Editors        string
	Viewers        string
	TrackingNumber string
}

type PostedFileTree struct {
	From           common.Address
	Account        string
	HashParent     string
	HashChild      string
	Contents       string
	Viewers        string
	Editors        string
	TrackingNumber string
}

type AddedViewers struct {
	From       common.Address
	ViewerIds  string
	ViewerKeys string
	ForAddress string
	FileOwner  string
}

type RemovedViewers struct {
	From       common.Address
	ViewerIds  string
	ForAddress string
	FileOwner  string
}

type ResetViewers struct {
	From       common.Address
	ForAddress string
	FileOwner  string
}

type ChangedOwner struct {
	From       common.Address
	ForAddress string
	FileOwner  string
	NewOwner   string
}

type AddedEditors struct {
	From       common.Address
	EditorIds  string
	EditorKeys string
	ForAddress string
	FileOwner  string
}

type RemovedEditors struct {
	From       common.Address
	EditorIds  string
	ForAddress string
	FileOwner  string
}

type ResetEditors struct {
	From       common.Address
	ForAddress string
	FileOwner  string
}

type CreatedNotification struct {
	From            common.Address
	To              string
	Contents        string
	PrivateContents string
}

type DeletedNotification struct {
	From             common.Address
	NotificationFrom string
	Time             uint64
}

type BlockedSenders struct {
	From    common.Address
	ToBlock []string
}
