package main

import (
	"PBL_Proj/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to our page")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllTransactions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllTransactions")
	json.NewEncoder(w).Encode(models.Transactions)
}

func returnParticularTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	uintKey, _ := strconv.ParseUint(key, 10, 32)
	fmt.Println(uintKey, key, vars)
	for _, transaction := range models.Transactions {
		if transaction.TransactionId == uint32(uintKey) {
			json.NewEncoder(w).Encode(transaction)
		}
	}
}

func returnTransactionsOfParticularAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	uintKey, _ := strconv.ParseUint(key, 10, 32)
	fmt.Println(uintKey, key, vars)
	for _, transaction := range models.Transactions {
		if transaction.AccountId == uint32(uintKey) {
			json.NewEncoder(w).Encode(transaction)
		}
	}
}

func postTransaction(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var requestMap map[string]interface{}
	json.Unmarshal(reqBody, &requestMap)

	if !validatePostTransaction(w, requestMap) {
		return
	}
	var newTransaction models.Transaction
	json.Unmarshal(reqBody, &newTransaction)

	newTransaction.TimeStamp = time.Now().UTC()
	newTransaction.TransactionId = rand.Uint32()

	if !newTransaction.PurchaseType.IsValidPurchaseString() {
		http.Error(w, "Invalid Purchase Type, please mention a correct purchase string", http.StatusBadRequest)
		return
	} else if !newTransaction.TransactionType.IsValidTransactionString() {
		http.Error(w, "Invalid Transaction Type, Please mention correct transaction string", http.StatusBadRequest)
		return
	}

	models.Transactions = append(models.Transactions, newTransaction)

	json.NewEncoder(w).Encode(newTransaction)
}

func deleteTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	uintKey, _ := strconv.ParseUint(key, 10, 32)
	fmt.Println(uintKey, key, vars)
	for index, transaction := range models.Transactions {
		if transaction.TransactionId == uint32(uintKey) {
			models.Transactions = append(models.Transactions[:index], models.Transactions[index+1:]...)
			fmt.Fprintf(w, "Successfully deleted Transaction id no: %s\n", key)
			json.NewEncoder(w).Encode(transaction)
			break
		}
	}
}

func putTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	uintKey, _ := strconv.ParseUint(key, 10, 32)

	reqBody, _ := ioutil.ReadAll(r.Body)
	var requestMap map[string]interface{}

	json.Unmarshal(reqBody, &requestMap)
	errorPresent, columnPresent := validatePutTransaction(w, requestMap)
	if errorPresent {
		return
	}

	var newTransaction models.Transaction
	json.Unmarshal(reqBody, &newTransaction)

	for idx, transaction := range models.Transactions {
		if transaction.TransactionId == uint32(uintKey) {
			if columnPresent["isCredit"] {
				transaction.IsCredit = newTransaction.IsCredit
			}
			if columnPresent["accountId"] {
				transaction.AccountId = newTransaction.AccountId
			}
			if columnPresent["transactionType"] {
				transaction.TransactionType = newTransaction.TransactionType
			}
			if columnPresent["purchaseType"] {
				transaction.PurchaseType = newTransaction.PurchaseType
			}
			if columnPresent["amount"] {
				transaction.Amount = newTransaction.Amount
			}
			json.NewEncoder(w).Encode(transaction)
			models.Transactions[idx] = transaction
			break
		}
	}
}

func returnTransactionAnalysisOfAnAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	uintKey, _ := strconv.ParseUint(key, 10, 32)

	reqBody, _ := ioutil.ReadAll(r.Body)
	var requestMap map[string]interface{}

	json.Unmarshal(reqBody, &requestMap)
	errorPresent, columnPresent := validateGetCombinationTransaction(w, requestMap)
	if errorPresent {
		return
	}

	var transactionRequested models.Transaction
	json.Unmarshal(reqBody, &transactionRequested)

	var satisfyingTransactions []models.Transaction
	var sum float32 = 0.0
	for _, transaction := range models.Transactions {
		if transaction.AccountId == uint32(uintKey) {
			if (!columnPresent["isCredit"] || transactionRequested.IsCredit == transaction.IsCredit) &&
				(!columnPresent["transactionType"] || transactionRequested.TransactionType == transaction.TransactionType) &&
				(!columnPresent["purchaseType"] || transactionRequested.PurchaseType == transaction.PurchaseType) &&
				((!columnPresent["amount"] || (transactionRequested.Amount > 0 && transactionRequested.Amount < transaction.Amount)) ||
					(!columnPresent["amount"] || (transactionRequested.Amount < 0 && -transactionRequested.Amount > transaction.Amount))) {
				sum += transaction.Amount
				satisfyingTransactions = append(satisfyingTransactions, transaction)
			}

		}
	}
	json.NewEncoder(w).Encode(satisfyingTransactions)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/allTransactions", returnAllTransactions)
	myRouter.HandleFunc("/transaction/{id}", returnParticularTransaction).Methods("GET")
	myRouter.HandleFunc("/allTransactionsOfAccount/{id}", returnTransactionsOfParticularAccount).Methods("GET")
	myRouter.HandleFunc("/analyzeTransactionsOfAccount/{id}", returnTransactionAnalysisOfAnAccount).Methods("GET")

	myRouter.HandleFunc("/transaction", postTransaction).Methods("POST")
	myRouter.HandleFunc("/transaction/{id}", deleteTransaction).Methods("DELETE")
	myRouter.HandleFunc("/transaction/{id}", putTransaction).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":10000", myRouter))

}

func main() {
	var AccountId uint32 = 450012
	var AccountIdB uint32 = 450015
	models.Transactions = []models.Transaction{
		{TimeStamp: time.Now().UTC(), IsCredit: models.Debit, TransactionId: rand.Uint32(), AccountId: AccountId, TransactionType: models.DebitCard, PurchaseType: models.Ecommerce, Amount: float32(32.55)},
		{TimeStamp: time.Now().UTC(), IsCredit: models.Debit, TransactionId: rand.Uint32(), AccountId: AccountIdB, TransactionType: models.CreditCard, PurchaseType: models.Grocery, Amount: float32(178.20)},
		{TimeStamp: time.Now().UTC(), IsCredit: models.Debit, TransactionId: rand.Uint32(), AccountId: AccountId, TransactionType: models.Cash, PurchaseType: models.Grocery, Amount: float32(6000.00)},
		{TimeStamp: time.Now().UTC(), IsCredit: models.Debit, TransactionId: rand.Uint32(), AccountId: AccountIdB, TransactionType: models.CreditCard, PurchaseType: models.Fitness, Amount: float32(5663.00)},
		{TimeStamp: time.Now().UTC(), IsCredit: models.Debit, TransactionId: rand.Uint32(), AccountId: AccountId, TransactionType: models.MobileApplication, PurchaseType: models.Investment, Amount: float32(2345.00)},
		{TimeStamp: time.Now().UTC(), IsCredit: models.Credit, TransactionId: rand.Uint32(), AccountId: AccountId, TransactionType: models.Cash, PurchaseType: models.None, Amount: float32(10000.00)},
	}
	handleRequests()
}
