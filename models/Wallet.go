package models

type Wallet struct {
	WalletID        uint64  `json:"id_wallet,omitempty"`
	UserName        string  `json:"user_name,omitempty"`
	AccountID       uint64  `json:"id_account,omitempty"`
	Current_balance float64 `json:"current_balance,omitempty"`
}
