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

`POST` `http://localhost:8080/books` insert a new book with informations passed into, for example:
```
{
  "name": "Game Of Thrones",
  "edition: "1",
  "publicationYear": "1996",
  "authors": "George R. R. Martin"
}
```

`PATCH` `http://localhost:8080/books/{id}` patch the book of id that you specified in URL param with the informations.
Data without updates:
```
{
  "id": 1,
  "name": "Game Of Thrones",
  "edition: "1",
  "publicationYear": "1996",
  "authors": "George R. R. Martin"
}
```

Data with updates:
```
{
  "id": 1,
  "name": "Game Of Thrones",
  "edition: "1",
  "publicationYear": "1997",
  "authors": "George R. R. Martin"
}
```
`DELETE` `http://localhost:8080/books/{id}` delete the book of id that you specified in URL param.

