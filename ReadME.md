# Build a build an API service using Extend APIs

1. List all virtual cards available to your user, including the available balance remaining. In this case, we would expect to see just one virtual card returned.

2. List the transactions associated with your virtual card.

3. View the details for each individual transaction you’ve made.

4. API responses contain many fields. For each of the endpoints you expose, please return a “lite” view picking just a few important fields that demonstrate the main pieces of each response.

Node is a great choice for building command line tools.
In this project, you see how to build a Node CLI in-memory database.

## Requirements
* [Go](https://go.dev/)

## Installation and Execution (if applicable)

1. $ git clone git@github.com:shambhavidinakar/extend.git
2. $ cd extend
3. $ go build server.go 
4. $./server

## APIs

Base URL : http://localhost:8080
The list of users APIs:

|METHOD|URL|PARAMETERS|DESCRIPTION|
|------|---|---------------|----------------------------------|
|POST|/login |email(string),password(string) | Login with email and password |
|GET|/getAllCards |lite(optional) | Get all cards of the user(logged in) |
|GET|/getAllTransactionsForCard | id(string),lite(optional)| Get all transactions of given card id |
|GET|/getTransactionDetails | id(string),lite(optional) | Get all transaction details of given transaction id |
|GET|/getAllCardsTransactions | lite(optional) | Get all cards and all transactions of each card of the user(logged in)|

