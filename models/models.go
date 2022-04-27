package models

import (
	"time"
)

type creditOrDebit bool

const (
	Credit creditOrDebit = true
	Debit  creditOrDebit = false
)

type transactionString string

const (
	DebitCard         transactionString = "Debit Card"
	CreditCard        transactionString = "Credit Card"
	Cash              transactionString = "Cash"
	OnlineBanking     transactionString = "Online Banking"
	MobileApplication transactionString = "Mobile Application"
)

func (k transactionString) IsValidTransactionString() bool {
	switch k {
	case DebitCard:
	case CreditCard:
	case Cash:
	case OnlineBanking:
	case MobileApplication:
	default:
		return false
	}
	return true
}

type purchaseString string

const (
	Ecommerce  purchaseString = "E-Commerce"
	Grocery    purchaseString = "Grocery"
	Fitness    purchaseString = "Fitness"
	Transport  purchaseString = "Transport"
	Luxury     purchaseString = "Luxury"
	Investment purchaseString = "Investment"
	None       purchaseString = "None"
)

func (k purchaseString) IsValidPurchaseString() bool {
	switch k {
	case Ecommerce:
	case Grocery:
	case Fitness:
	case Transport:
	case Luxury:
	case Investment:
	case None:
	default:
		return false
	}
	return true
}

type Transaction struct {
	TimeStamp       time.Time         `json:"timeStamp"`
	IsCredit        creditOrDebit     `json:"isCredit"`
	TransactionId   uint32            `json:"transactionId"`
	AccountId       uint32            `json:"accountId"`
	TransactionType transactionString `json:"transactionType"`
	PurchaseType    purchaseString    `json:"purchaseType"`
	Amount          float32           `json:"amount"`
}

var Transactions []Transaction

func main() {
}
