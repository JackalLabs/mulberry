package types

type ExecuteMsg struct {
	PostKey           *ExecuteMsgPostKey           `json:"post_key,omitempty"`
	PostFile          *ExecuteMsgPostFile          `json:"post_file,omitempty"`
	BuyStorage        *ExecuteMsgBuyStorage        `json:"buy_storage,omitempty"`
	DeleteFile        *ExecuteMsgDeleteFile        `json:"delete_file,omitempty"`
	RequestReportForm *ExecuteMsgRequestReportForm `json:"request_report_form,omitempty"`
	DeleteFileTree    *ExecuteMsgDeleteFileTree    `json:"delete_file_tree,omitempty"`
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
	ForAddress   string `json:"for_address,omitempty"`
	DurationDays int64  `json:"duration_days,omitempty"`
	Bytes        int64  `json:"bytes,omitempty"`
	PaymentDenom string `json:"payment_denom,omitempty"`
	Referral     string `json:"referral,omitempty"`
}

type ExecuteMsgDeleteFile struct {
	Merkle string `json:"merkle,omitempty"`
	Start  int64  `json:"start,omitempty"`
}

type ExecuteMsgRequestReportForm struct {
	Prover string `json:"prover,omitempty"`
	Merkle string `json:"merkle,omitempty"`
	Owner  string `json:"owner,omitempty"`
	Start  int64  `json:"start,omitempty"`
}

type ExecuteMsgDeleteFileTree struct {
	HashPath string `json:"hash_path,omitempty"`
	Account  string `json:"account,omitempty"`
}

// ToString returns a string representation of the message
func (m *ExecuteMsg) ToString() string {
	return toString(m)
}
