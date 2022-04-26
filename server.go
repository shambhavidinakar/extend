package main

import (
	"encoding/json"
	"extend/extendReqs"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var access_token string

const (
	port = "8080"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := "Unable to Authenticate"
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		email := r.FormValue("email")
		password := r.FormValue("password")
		req := extendReqs.Req{}
		result := req.Login(email, password)

		if result == "200 OK" {
			message = "Logged in Successfully"
		}
	}
	json.NewEncoder(w).Encode(message)
}

func CardsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		lite := r.FormValue("lite")
		req := extendReqs.Req{}
		result := req.GetAllCards(lite)
		json.NewEncoder(w).Encode(result)
	}
}

func TransactionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		id := r.FormValue("id")
		lite := r.FormValue("lite")
		req := extendReqs.Req{}
		result := req.GetTransactions(id, lite)
		json.NewEncoder(w).Encode(result)
	}
}

func TransactionDetailsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		id := r.FormValue("id")
		lite := r.FormValue("lite")
		req := extendReqs.Req{}
		result := req.GetTransactionDetails(id, lite)
		json.NewEncoder(w).Encode(result)
	}
}

func CardsTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		lite := r.FormValue("lite")
		req := extendReqs.Req{}
		result := req.GetAllCardsTransactions(lite)
		json.NewEncoder(w).Encode(result)
	}
}

func main() {
	fmt.Println("Starting web server...")
	r := mux.NewRouter()
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/getAllCards", CardsHandler)
	r.HandleFunc("/getAllTransactionsForCard", TransactionsHandler)
	r.HandleFunc("/getTransactionDetails", TransactionDetailsHandler)
	r.HandleFunc("/getAllCardsTransactions", CardsTransactions)
	http.ListenAndServe(":"+port, r)
}
