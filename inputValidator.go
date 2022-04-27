package main

import (
	"net/http"
)

func validatePostTransaction(w http.ResponseWriter, request map[string]interface{}) bool {
	if _, ok := request["isCredit"]; !ok {
		http.Error(w, "[ERROR] 'isCredit' field missing", http.StatusBadRequest)
		return false
	}
	if _, ok := request["accountId"]; !ok {
		http.Error(w, "[ERROR] 'accountId' field missing", http.StatusBadRequest)
		return false
	}
	if _, ok := request["transactionType"]; !ok {
		http.Error(w, "[ERROR] 'transactionType' field missing", http.StatusBadRequest)
		return false
	}
	if _, ok := request["purchaseType"]; !ok {
		http.Error(w, "[ERROR] 'purchaseType' field missing", http.StatusBadRequest)
		return false
	}
	if _, ok := request["amount"]; !ok {
		http.Error(w, "[ERROR] 'amount' field missing", http.StatusBadRequest)
		return false
	}
	if !IsaNumber(request["amount"]) {
		http.Error(w, "[ERROR] 'amount' not a number or floating point", http.StatusBadRequest)
		return false
	}
	return true
}

func validatePutTransaction(w http.ResponseWriter, request map[string]interface{}) (bool, map[string]bool) {
	_, isCredit := request["isCredit"]
	_, accountId := request["accountId"]
	_, transactionType := request["transactionType"]
	_, purchaseType := request["purchaseType"]
	_, amount := request["amount"]

	isPresentArray := map[string]bool{
		"isCredit":        isCredit,
		"accountId":       accountId,
		"transactionType": transactionType,
		"purchaseType":    purchaseType,
		"amount":          amount,
	}
	var k bool = false
	for _, value := range isPresentArray {
		k = k || value
	}
	if !k {
		http.Error(w, "[WARNING] There is no field to update", http.StatusOK)
		return true, isPresentArray
	}
	if amount && !IsaNumber(request["amount"]) {
		http.Error(w, "[ERROR] 'amount' not a number or floating point", http.StatusBadRequest)
		return true, isPresentArray
	}
	return false, isPresentArray
}

func validateGetCombinationTransaction(w http.ResponseWriter, request map[string]interface{}) (bool, map[string]bool) {
	_, isCredit := request["isCredit"]
	_, accountId := request["accountId"]
	_, transactionType := request["transactionType"]
	_, purchaseType := request["purchaseType"]
	_, amount := request["amount"]

	isPresentArray := map[string]bool{
		"isCredit":        isCredit,
		"accountId":       accountId,
		"transactionType": transactionType,
		"purchaseType":    purchaseType,
		"amount":          amount,
	}
	if amount && !IsaNumber(request["amount"]) {
		http.Error(w, "[ERROR] 'amount' not a number or floating point", http.StatusBadRequest)
		return true, isPresentArray
	}
	return false, isPresentArray
}
