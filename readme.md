# Simple Library Application with API and PostgreSQL Database

This is a basic library application that provides an API for managing book details. It allows users to add new books to the library and retrieve a list of all books stored in the PostgreSQL database.

## Requirements

- Go programming language (version 1.20)
- PostgreSQL database
- `github.com/lib/pq` package for PostgreSQL driver (automatically installed during setup)

## Setup

1. Clone the repository or download the source code:

   ```bash
   git clone https://github.com/krunalhinguu/citadel-go
   cd citadel/assignment
   ```

2. Install the required dependencies:

    ``` 
    go mod tidy && go mod vendor
    ```

3. Update the database credentials in main.go

4. Run the application :

    ```
    go run main.go
    ```

## Endpoints

#### Add a new book
- To add a new book, make a POST request to /books with the book details in JSON format.


Request:


``` 
curl -X POST -H "Content-Type: application/json" -d '{"title":"Book 1", "author":"Author 1"}' http://localhost:8080/books
Response (HTTP status code 201 - Created):
```
Response

```
{"status":"Created"}
```

#### Get all books
- To retrieve a list of all books, make a GET request to /books.

Request:

```
curl http://localhost:8080/books
```


Response :
```
[
    {"id":1,"title":"Book 1","author":"Author 1"},
    {"id":2,"title":"Book 2","author":"Author 2"},
    {"id":3,"title":"Book 3","author":"Author 3"}
]
```