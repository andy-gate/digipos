package models

type HistoryPurchase struct {
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
}

type HistoryDeposit struct {
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