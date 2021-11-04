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

type Response struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Meta struct {
		Limit int `json:"limit"`
		Page int `json:"page"`
		TotalPage int `json:"totalPage"`
		TotalRecord int `json:"totalRecord"`
	} `json:"meta"`
	Data [] struct {
		Id int `json:"id"`
		LinkAjaNo string `json:"linkAjaNo"`
		InitiationTime string `json:"initiationTime"`
		ServiceName string `json:"serviceName"`
		InitiatorParty  string `json:"initiatorParty"`
		CreditParty string `json:"creditParty"`
		DebitParty  string `json:"debitParty"`
		TransactionStatus  string `json:"transactionStatus"`
		TransactionAmount int `json:"transactionAmount"`
		ReceiptNo string `json:"receiptNo"`
		TransactionFlag string `json:"transactionFlag"`
	} `json:"data"`
}