package types

type ExecuteMsg struct {
	PostKey            *ExecuteMsgPostKey            `json:"post_key,omitempty"`
	PostFile           *ExecuteMsgPostFile           `json:"post_file,omitempty"`
	BuyStorage         *ExecuteMsgBuyStorage         `json:"buy_storage,omitempty"`
	DeleteFile         *ExecuteMsgDeleteFile         `json:"delete_file,omitempty"`
	RequestReportForm  *ExecuteMsgRequestReportForm  `json:"request_report_form,omitempty"`
	DeleteFileTree     *ExecuteMsgDeleteFileTree     `json:"delete_file_tree,omitempty"`
	ProvisionFileTree  *ExecuteMsgProvisionFileTree  `json:"provision_file_tree,omitempty"`
	PostFileTree       *ExecuteMsgPostFileTree       `json:"post_file_tree,omitempty"`
	AddViewers         *ExecuteMsgAddViewers         `json:"add_viewers,omitempty"`
	RemoveViewers      *ExecuteMsgRemoveViewers      `json:"remove_viewers,omitempty"`
	ResetViewers       *ExecuteMsgResetViewers       `json:"reset_viewers,omitempty"`
	ChangeOwner        *ExecuteMsgChangeOwner        `json:"change_owner,omitempty"`
	AddEditors         *ExecuteMsgAddEditors         `json:"add_editors,omitempty"`
	RemoveEditors      *ExecuteMsgRemoveEditors      `json:"remove_editors,omitempty"`
	ResetEditors       *ExecuteMsgResetEditors       `json:"reset_editors,omitempty"`
	CreateNotification *ExecuteMsgCreateNotification `json:"create_notification,omitempty"`
	DeleteNotification *ExecuteMsgDeleteNotification `json:"delete_notification,omitempty"`
	BlockedSenders     *ExecuteMsgBlockedSenders     `json:"block_senders,omitempty"`
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

type ExecuteMsgAddViewers struct {
	ViewerIds  string `json:"viewer_ids"`
	ViewerKeys string `json:"viewer_keys"`
	Address    string `json:"address"`
	FileOwner  string `json:"file_owner"`
}

type ExecuteMsgRemoveViewers struct {
	ViewerIds string `json:"viewer_ids"`
	Address   string `json:"address"`
	FileOwner string `json:"file_owner"`
}

type ExecuteMsgResetViewers struct {
	Address   string `json:"address"`
	FileOwner string `json:"file_owner"`
}

type ExecuteMsgChangeOwner struct {
	Address   string `json:"address"`
	FileOwner string `json:"file_owner"`
	NewOwner  string `json:"new_owner"`
}

type ExecuteMsgAddEditors struct {
	EditorIds  string `json:"editor_ids"`
	EditorKeys string `json:"editor_keys"`
	Address    string `json:"address"`
	FileOwner  string `json:"file_owner"`
}

type ExecuteMsgRemoveEditors struct {
	EditorIds string `json:"editor_ids"`
	Address   string `json:"address"`
	FileOwner string `json:"file_owner"`
}

type ExecuteMsgResetEditors struct {
	Address   string `json:"address"`
	FileOwner string `json:"file_owner"`
}

type ExecuteMsgCreateNotification struct {
	To              string `json:"to"`
	Contents        string `json:"contents"`
	PrivateContents string `json:"private_contents"`
}

type ExecuteMsgDeleteNotification struct {
	From string `json:"from"`
	Time int64  `json:"time"`
}

type ExecuteMsgBlockedSenders struct {
	ToBlock []string `json:"to_block"`
}

// ToString returns a string representation of the message
func (m *ExecuteMsg) ToString() string {
	return toString(m)
}
