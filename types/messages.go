package types

type ExecuteMsg struct {
	PostKey           *ExecuteMsgPostKey           `json:"post_key,omitempty"`
	PostFile          *ExecuteMsgPostFile          `json:"post_file,omitempty"`
	BuyStorage        *ExecuteMsgBuyStorage        `json:"buy_storage,omitempty"`
	DeleteFile        *ExecuteMsgDeleteFile        `json:"delete_file,omitempty"`
	RequestReportForm *ExecuteMsgRequestReportForm `json:"request_report_form,omitempty"`
	DeleteFileTree    *ExecuteMsgDeleteFileTree    `json:"delete_file_tree,omitempty"`
	ProvisionFileTree *ExecuteMsgProvisionFileTree `json:"provision_file_tree,omitempty"`
	PostFileTree      *ExecuteMsgPostFileTree      `json:"post_file_tree,omitempty"`
}

type ExecuteMsgPostKey struct {
	Key string `json:"key"`
}

type ExecuteMsgPostFile struct {
	Merkle        string `json:"merkle"`
	FileSize      int64  `json:"file_size"`
	ProofInterval int64  `json:"proof_interval"`
	ProofType     int64  `json:"proof_type"`
	MaxProofs     int64  `json:"max_proofs"`
	Expires       int64  `json:"expires"`
	Note          string `json:"note"`
}

type ExecuteMsgBuyStorage struct {
	ForAddress   string `json:"for_address"`
	DurationDays int64  `json:"duration_days"`
	Bytes        int64  `json:"bytes"`
	PaymentDenom string `json:"payment_denom"`
	Referral     string `json:"referral"`
}

type ExecuteMsgDeleteFile struct {
	Merkle string `json:"merkle"`
	Start  int64  `json:"start"`
}

type ExecuteMsgRequestReportForm struct {
	Prover string `json:"prover"`
	Merkle string `json:"merkle"`
	Owner  string `json:"owner"`
	Start  int64  `json:"start"`
}

type ExecuteMsgDeleteFileTree struct {
	HashPath string `json:"hash_path"`
	Account  string `json:"account"`
}

type ExecuteMsgProvisionFileTree struct {
	Editors        string `json:"editors"`
	Viewers        string `json:"viewers"`
	TrackingNumber string `json:"tracking_number"`
}

type ExecuteMsgPostFileTree struct {
	Account        string `json:"account"`
	HashParent     string `json:"hash_parent"`
	HashChild      string `json:"hash_child"`
	Contents       string `json:"contents"`
	Viewers        string `json:"viewers"`
	Editors        string `json:"editors"`
	TrackingNumber string `json:"tracking_number"`
}

// ToString returns a string representation of the message
func (m *ExecuteMsg) ToString() string {
	return toString(m)
}
