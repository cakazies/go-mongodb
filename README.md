# GO-POSTGRESQL
Golang RESTFUL API with Database MongoDB (NOSQL)
and with ECHO framework, validate

## Collection Postman
you can import collection in `test/postman/go-mongo.postman_collection`


##  How to run
- you can install mongodb in your computer tutorial this [link](https://docs.mongodb.com/v3.2/tutorial/install-mongodb-on-windows/) 
- install dependencies go get
    - `go get github.com/globalsign/mgo/bson`
    - `go get go.mongodb.org/mongo-driver/bson`
    - `go get github.com/labstack/echo`
    - `go get go.mongodb.org/mongo-driver/mongo`

## Fitur
about all this fitur you can read this repo in WIKI or click this [link](https://github.com/cakazies/go-mongodb/wiki)

## Unit Testing
run Testing with this command
> go test ./test

#### Run Local
`go run main.go`

## Run Docker
- build docker `docker build -t go-mongo-images .`
- Run `docker run -it --rm --name cont-go-mongo go-mongo-images`

