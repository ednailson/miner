# Miner

Miner implements a Stratum V1 server that stores the requests on a database.

## Application

### Requirements 

* [docker-compose](https://docs.docker.com/compose/)

### Starting app

To start the app you can simply run the command below

> make start

The server will run at port `4000` and the database at port `5432`. 

Make sure that those ports are available on your machine.

### Client

There is a simple [client](examples/client/client.go) that can be used to request to the server. 

Be sure that the server has started running the [server](README.md#starting-app).

## Tests

### Requirements

* [golang](https://golang.org/)

### Running

To run the tests you can simply run the command bellow

> make tests

# Developer

Ednailson Junior

* evbcjr@gmail.com
* [LinkedIn](https://www.linkedin.com/in/ednailsonvb/)
* [github](https://github.com/ednailson)
