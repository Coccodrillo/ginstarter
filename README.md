
#Gin API starter

This is just a simple starter project in Golang with [gin](https://github.com/gin-gonic/gin) as HTTP framework and [gorp](https://github.com/go-gorp/gorp) to manage Database.

I used a similar setup in 100k plus user production environment with success.

To install do the following - Import .sql file into a database and adjust db access params in the .env file

then

```go
go get github.com/Coccodrillo/ginstarter
```
```go
go get .
```
```go
go build
```

It's missing all the bells and whistles like validation, graceful restart, convenience methods and the like, because it's meant to illustrate a simple usage example for the framework.