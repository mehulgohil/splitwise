package models

type TransactionModel struct {
	TransactionId int `json:"-"`
	TransactionPerPlaceId int `json:"-"`
	TotalAmount int    `json:"totalAmount"`
	Place       string `json:"place"`
	Date        string `json:"date"`
	SpentBy     struct {
		PersonID int `json:"-"`
		Mobile string `json:"mobile"`
		Name   string `json:"name"`
	} `json:"spentBy"`
	NPeople int `json:"nPeople,omitempty"`
	Split   []PostTransactionRequestSplit `json:"split,omitempty"`
	MyShare       int         `json:"myShare,omitempty"`
	PaymentStatus string      `json:"paymentStatus,omitempty"`
}

type PostTransactionRequestSplit struct {
	PersonID int `json:"-"`
	Mobile      string `json:"mobile"`
	Name        string `json:"name"`
	ShareAmount int    `json:"shareAmount"`
	PaymentStatus string `json:"paymentStatus"`
}

type PatchTransactionRequest struct {
	Mobile string `json:"mobile"`
}

