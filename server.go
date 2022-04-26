package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var access_token string

type Auth struct {
	Token string `json:"token"`
}

type CardHolder struct {
	Id               string `json:"id"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Email            string `json:"email"`
	OrganizationID   string `json:"organizationId"`
	OrganizationRole string `json:"organizationRole"`
}

type Pagination struct {
	Page          int `json:"page"`
	PageItemCount int `json:"pageItemCount"`
	TotalItems    int `json:"totalItems"`
	NoOfPages     int `json:"numberOfPages"`
}

type VirtualCard struct {
	Id                 string     `json:"id"`
	Status             string     `json:"status"`
	CardHolder         CardHolder `json:"cardholder"`
	Name               string     `json:"displayName"`
	Expires            string     `json:"expires"`
	Currency           string     `json:"currency"`
	LimitCents         int        `json:"limitCents"`
	BalanceCents       int        `json:"balanceCents"`
	SpendCents         int        `json:"spentCents"`
	LifeTimeSpentCents int        `json:"lifetimeSpentCents"`
	Network            string     `json:"network"`
	CompanyName        string     `json:"companyName"`
}

type VirtualCards struct {
	Pagination   Pagination    `json:"pagination"`
	VirtualCards []VirtualCard `json:"virtualCards"`
}

type Transaction struct {
	Id                         string    `json:"id"`
	CardHolderName             string    `json:"cardholderName"`
	VCDisplayName              string    `json:"vcnDisplayName"`
	VirtualCardId              string    `json:"virtualCardId"`
	Type                       string    `json:"type"`
	Status                     string    `json:"status"`
	ApprovalCode               string    `json:"approvalCode"`
	AuthBillingAmountCents     int       `json:"authBillingAmountCents"`
	AuthBillingCurrency        string    `json:"authBillingCurrency"`
	ClearingBillingAmountCents int       `json:"clearingBillingAmountCents"`
	ClearingBillingCurrency    string    `json:"clearingBillingCurrency"`
	MccGroup                   string    `json:"mccGroup"`
	MerchentId                 string    `json:"merchantId"`
	MerchentName               string    `json:"merchantName"`
	AuthedAt                   time.Time `json:"authedAt"`
	ClearedAt                  time.Time `json:"clearedAt"`
}

type Transactions struct {
	Transactions []Transaction `json:"transactions"`
}

type ResultLite struct {
	Card         VirtualCard   `json:"virtualCards"`
	Transactions []Transaction `json:"transactions"`
}

type Result struct {
	Card         interface{}   `json:"virtualCards"`
	Transactions []interface{} `json:"transactions"`
}

const (
	port = "8080"
)

func login(email string, password string) {
	httpposturl := "https://api.paywithextend.com/signin"
	values := map[string]string{"email": email, "password": password}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}
	request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(json_data))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/vnd.paywithextend.v2021-03-12+json")
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	var auth Auth
	err2 := json.Unmarshal(body, &auth)
	if err2 != nil {
		log.Fatal(err)
	}
	access_token = auth.Token
}

func getAllCards(lite string) []interface{} {
	httpposturl := "https://api.paywithextend.com/virtualcards"
	request, error := http.NewRequest("GET", httpposturl, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("Accept", "application/vnd.paywithextend.v2021-03-12+json")
	request.Header.Add("Authorization", "Bearer "+access_token)
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	if lite == "" {
		var res map[string]interface{}
		err := json.Unmarshal(body, &res)
		if err != nil {
			log.Fatal(err)
		}
		cards := res["virtualCards"].([]interface{})
		return cards
	} else {
		var res VirtualCards
		err := json.Unmarshal(body, &res)
		if err != nil {
			log.Fatal(err)
		}
		b, _ := json.Marshal(&res.VirtualCards)
		var cards []interface{}
		err = json.Unmarshal(b, &cards)
		return cards
	}
}

func getTransactions(id string, lite string) []interface{} {
	httpposturl := "https://api.paywithextend.com/virtualcards/" + id + "/transactions"
	request, error := http.NewRequest("GET", httpposturl, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("Accept", "application/vnd.paywithextend.v2021-03-12+json")
	request.Header.Add("Authorization", "Bearer "+access_token)
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	if lite == "" {
		var res map[string]interface{}
		err := json.Unmarshal(body, &res)
		if err != nil {
			log.Fatal(err)
		}
		transactions := res["transactions"].([]interface{})
		return transactions
	} else {
		var res Transactions
		err := json.Unmarshal(body, &res)
		if err != nil {
			log.Fatal(err)
		}
		b, _ := json.Marshal(&res.Transactions)
		var transactions []interface{}
		err = json.Unmarshal(b, &transactions)
		if err != nil {
			log.Fatal(err)
		}
		return transactions
	}
}

func getTransactionDetails(id string, lite string) map[string]interface{} {
	httpposturl := "https://api.paywithextend.com/transactions" + id
	request, error := http.NewRequest("GET", httpposturl, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("Accept", "application/vnd.paywithextend.v2021-03-12+json")
	request.Header.Add("Authorization", "Bearer "+access_token)
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	if lite == "" {
		var res map[string]interface{}
		err := json.Unmarshal(body, &res)
		if err != nil {
			log.Fatal(err)
		}
		return res
	} else {
		var res Transaction
		err := json.Unmarshal(body, &res)
		if err != nil {
			log.Fatal(err)
		}
		b, _ := json.Marshal(&res)
		var m map[string]interface{}
		err = json.Unmarshal(b, &m)
		if err != nil {
			log.Fatal(err)
		}
		return m
	}
}

func getAllCardsTransactions(lite string) []interface{} {
	virtualCards := getAllCards(lite)
	var results_lite []ResultLite
	var results []Result
	for _, card := range virtualCards {
		c, _ := json.Marshal(&card)
		var cardPlaceholder VirtualCard
		err := json.Unmarshal(c, &cardPlaceholder)
		if err != nil {
			log.Fatal(err)
		}
		transactions := getTransactions(cardPlaceholder.Id, lite)
		t, _ := json.Marshal(&transactions)
		var transactionsPlaceholder []Transaction
		err = json.Unmarshal(t, &transactionsPlaceholder)
		if err != nil {
			log.Fatal(err)
		}
		if lite == "" {
			var result Result
			result.Card = card
			result.Transactions = transactions
			results = append(results, result)
		} else {
			var resultLite ResultLite
			resultLite.Card = cardPlaceholder
			resultLite.Transactions = transactionsPlaceholder
			results_lite = append(results_lite, resultLite)
		}
	}
	var finalresults []interface{}
	if lite == "" {
		r1, _ := json.Marshal(&results)
		err := json.Unmarshal(r1, &finalresults)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		r2, _ := json.Marshal(&results_lite)
		err := json.Unmarshal(r2, &finalresults)
		if err != nil {
			log.Fatal(err)
		}
	}
	return finalresults
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		email := r.FormValue("email")
		password := r.FormValue("password")
		login(email, password)
	}
}

func CardsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		lite := r.FormValue("lite")
		result := getAllCards(lite)
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
		result := getTransactions(id, lite)
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
		result := getTransactionDetails(id, lite)
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
		result := getAllCardsTransactions(lite)
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
