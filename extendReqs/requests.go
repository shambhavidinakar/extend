package extendReqs

import (
	"bytes"
	"encoding/json"
	"extend/models"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Req struct{}

var access_token string

func (Req) Login(email string, password string) string {
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
	auth := models.Auth{}
	err2 := json.Unmarshal(body, &auth)
	if err2 != nil {
		log.Fatal(err)
	}
	access_token = auth.Token
	return response.Status
}

func (Req) GetAllCards(lite string) []interface{} {
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
		res := models.VirtualCards{}
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

func (Req) GetTransactions(id string, lite string) []interface{} {
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
		res := models.Transactions{}
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

func (Req) GetTransactionDetails(id string, lite string) map[string]interface{} {
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
		res := models.Transaction{}
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

func (r Req) GetAllCardsTransactions(lite string) []interface{} {
	virtualCards := r.GetAllCards(lite)
	results_lite := []models.ResultLite{}
	results := []models.Result{}
	for _, card := range virtualCards {
		c, _ := json.Marshal(&card)
		cardPlaceholder := models.VirtualCard{}
		err := json.Unmarshal(c, &cardPlaceholder)
		if err != nil {
			log.Fatal(err)
		}
		transactions := r.GetTransactions(cardPlaceholder.Id, lite)
		t, _ := json.Marshal(&transactions)
		transactionsPlaceholder := []models.Transaction{}
		err = json.Unmarshal(t, &transactionsPlaceholder)
		if err != nil {
			log.Fatal(err)
		}
		if lite == "" {
			result := models.Result{}
			result.Card = card
			result.Transactions = transactions
			results = append(results, result)
		} else {
			resultLite := models.ResultLite{}
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
