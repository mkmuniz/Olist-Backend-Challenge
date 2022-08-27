# Olist-Backend-Challenge

**First Step**

Clone the repository

**Second Step**

Open the folder and create a file called `config.toml` and add the following settings with filled fields:

```
[api]
port=""

[db]
url=""
name=""
```

**Third Step**

Type `go run main.go` and acess the API in `http://localhost:8080`

**Fourth Step**

You can acess `http://localhost:8080/books` and do the following requests:

`GET` `http://localhost:8080/books` return all books registered in database

`GET` `http://localhost:8080/books/{id}` return the book of id that you specified in URL param.

`POST` `http://localhost:8080/books` insert a new book with informations passed into.

`PATCH` `http://localhost:8080/books/{id}` patch the book of id that you specified in URL param with the informations.

`DELETE` `http://localhost:8080/books/{id}` delete the book of id that you specified in URL param.

