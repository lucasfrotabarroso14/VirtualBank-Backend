package models

type Transaction struct {
	IDtransaction        uint64  `json:"id_transaction, omitempty"`
	IDaccount            uint64  `json:"id_account, omitempty"`
	TransactionType      string  `json:"transaction_type, omitempty"`
	Amount               float64 `json:"amount, omitempty"`
	Description          string  `json:"description, omitempty"`
	Category             string  `json:"category, omitempty"`
	IDdestinationAccount uint64  `json:"id_destination_account, omitempty"`
}
