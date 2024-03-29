package models

type Transaction struct {
	IDtransaction          uint64  `json:"id_transaction,omitempty"`
	IDoriginAccount        uint64  `json:"id_origin_account,omitempty"`
	AccountOriginName      string  `json:"account_origin_name,omitempty"`
	TransactionType        string  `json:"transaction_type,omitempty"`
	Amount                 float64 `json:"amount,omitempty"`
	Description            string  `json:"description,omitempty"`
	Category               string  `json:"category,omitempty"`
	IDdestinationAccount   uint64  `json:"id_destination_account,omitempty"`
	AccountDestinationName string  `json:"account_destination_name,omitempty"`
}
