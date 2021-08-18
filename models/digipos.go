package models

type HistoryPurchase struct {
	LinkAjaNo string `json:"link_aja_no"`
	InitiationTime string `json:"initiation_time"`
	ServiceName string `json:"service_name"`
	InitiatorParty  string `json:"initiator_party"`
	CreditParty string `json:"credit_party"`
	DebitParty  string `json:"debit_party"`
	TransactionStatus  string `json:"transaction_status"`
	TransactionAmount int `json:"transaction_amount"`
	ReceiptNo string `json:"receipt_no"`
}

type HistoryDeposit struct {
	LinkAjaNo string `json:"link_aja_no"`
	InitiationTime string `json:"initiation_time"`
	ServiceName string `json:"service_name"`
	InitiatorParty  string `json:"initiator_party"`
	CreditParty string `json:"credit_party"`
	DebitParty  string `json:"debit_party"`
	TransactionStatus  string `json:"transaction_status"`
	TransactionAmount int `json:"transaction_amount"`
	ReceiptNo string `json:"receipt_no"`
}