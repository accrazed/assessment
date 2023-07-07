# Jackbox Assessment Server

A basic go/gin API for managing user data

## Data Stored

* First Name
* Last Name
* Email
* Username
* Password
* Extra Data ##

## TODOS

* Create a Signup -> Data Creation endpoint
* Create a Signed in -> modify data endpoint
* Create an Admin Account
  * All User Data Endpoint

## Running

To start the server, run `go run main.go`

You can `curl -X GET http://localhost:8080/users/7fa50761-e46d-49eb-bf8c-180dc3eacc2a` to get my user data :)
